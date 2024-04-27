package service

import (
	"log"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateSendOTPRequest(req *proto.SendOTPRequest) error {
	// Check if either email or phone is provided
	err := helpers.ValidateEmailOrPhone(req.GetEmail(), req.GetPhone())
	if err != nil {
		return err
	}
	// Validate action
	if !utils.IsSupportedSignupAction(req.GetAction()) {
		log.Println(req.GetAction())
		return status.Errorf(codes.InvalidArgument, "action must be either 'signup' or 'forget'")
	}
	return nil
}
