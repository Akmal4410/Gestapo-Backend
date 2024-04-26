package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (auth *authenticationService) SendOTP(ctx context.Context, req *proto.SendOTPRequest) (*proto.Response, error) {
	response := &proto.Response{}
	err := helpers.ValidateEmailOrPhone(req.GetEmail(), req.GetPhone())
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusBadRequest, "Invalid Email or Phone")
		return response, nil
	}

	column, value := helpers.IdentifiesColumnValue(req.Email, req.Phone)
	if req.Action == utils.SIGN_UP {
		if len(column) == 0 {
			auth.log.LogError("Error while IdentifiesColumnValue", column)
			response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
			return response, nil
		}
		res, err := auth.storage.CheckDataExist(column, value)
		if err != nil {
			auth.log.LogError("Error while CheckDataExist", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
			return response, nil
		}
		if res {
			err = fmt.Errorf("account already exist using this %s", column)
			auth.log.LogError(err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusNotFound, err.Error())
			return response, nil
		}
	}

	if !helpers.IsEmpty(req.Email) {
		err = auth.emailService.SendOTP(req.Email, utils.EmailSubject, utils.EmailSubject, auth.redis)
		if err != nil {
			auth.log.LogError("Error while SendOTP", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
			return response, nil
		}
	} else {
		phoneNumber := fmt.Sprintf("+91%s", req.Phone)
		err = auth.twilioService.SendOTP(phoneNumber)
		if err != nil {
			auth.log.LogError("Error while SendOTP", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
			return response, nil
		}
	}
	sessionToken, err := auth.token.CreateSessionToken(value, req.Action, time.Minute*5)
	if err != nil {
		auth.log.LogError("Error while CreateSessionToken", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return response, nil
	}

	response.StatusCode = http.StatusOK
	response.Status = true
	response.Message = "OTP sent successfully"
	response.ErrorInfo = nil

	mdOut := metadata.New(map[string]string{
		"session-token": sessionToken,
	})
	return response, grpc.SetHeader(ctx, mdOut)
}

func (auth *authenticationService) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.Response, error) {
	response := &proto.Response{}

	res, err := auth.storage.CheckDataExist("user_name", req.GetUserName())
	if err != nil {
		auth.log.LogError("Error while CheckDataExist", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return response, nil
	}
	if !res {
		err = fmt.Errorf("user doesnt exist %w", err)
		auth.log.LogError("User doesn't", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, "User doesnt exist")
		return response, nil
	}

	res, err = auth.storage.CheckPassword(req.UserName, req.Password)
	if err != nil {
		auth.log.LogError("Error while CheckPassword", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return response, nil
	}
	if !res {
		err = fmt.Errorf("wrong password %w", err)
		auth.log.LogError("Wrong password", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusForbidden, "User crediantials doesn't match")
		return response, nil
	}

	payload, err := auth.storage.GetTokenPayload("user_name", req.UserName)
	if err != nil {
		auth.log.LogError("Error while GetTokenPayload", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return response, nil
	}
	token, err := auth.token.CreateAccessToken(payload.UserId, req.UserName, payload.UserType, time.Hour*48)
	if err != nil {
		auth.log.LogError("Error while CreateAccessToken", err)
		response.ErrorInfo = helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return response, nil
	}
	response.StatusCode = http.StatusOK
	response.Status = true
	response.Message = "User loggedin Successfully"
	response.ErrorInfo = nil

	mdOut := metadata.New(map[string]string{
		"access-token": token,
	})
	return response, grpc.SetHeader(ctx, mdOut)
}
