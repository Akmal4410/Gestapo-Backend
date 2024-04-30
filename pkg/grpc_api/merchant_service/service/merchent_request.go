package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (merchant *merchantService) GetProfile(ctx context.Context, req *proto.GetMerchantProfileRequest) (*proto.GetMerchantProfileResponse, error) {
	if req.GetUserId() == "" {
		merchant.log.LogError("Error while Getting user id")
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}
	res, err := merchant.storage.CheckDataExist("user_data", "id", req.GetUserId())
	if err != nil {
		merchant.log.LogError("Error while CheckUserExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		err = fmt.Errorf("account does'nt exist using %s", req.GetUserId())
		merchant.log.LogError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	userData, err := merchant.storage.GetProfile(req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			merchant.log.LogError("Error while GetProfile", err)
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		merchant.log.LogError("Error while GetProfile", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if userData.ProfileImage != nil {
		url, err := merchant.s3.GetPreSignedURL(*userData.ProfileImage)
		if err != nil {
			merchant.log.LogError("Error while GetPreSignedURL", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
		userData.ProfileImage = &url
	}
	respone := &proto.GetMerchantProfileResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Profile fetched sucessfull",
		Data:    userData,
	}
	return respone, nil
}
