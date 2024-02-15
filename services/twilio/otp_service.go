package twilio

import (
	"fmt"

	"github.com/akmal4410/gestapo/utils"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	accountSid string
	authToken  string
	serviceSid string
)

type TwilioService interface {
	SendOTP(to string) error
	VerfiyOTP(to, code string) (bool, error)
}

type OTPService struct{}

func NewOTPService() TwilioService {
	return &OTPService{}
}

func LoadEnv() {
	accountSid = utils.EnvAccountSid()
	authToken = utils.EnvAuthToken()
	serviceSid = utils.EnvServiceSid()
}

func (service *OTPService) SendOTP(to string) error {
	LoadEnv()

	var client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	// params.SetCustomMessage("Your [Gestapo] verification code is:\n")

	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		return err
	}

	fmt.Printf("Verification has been send, Id :'%s'\n", *resp.AccountSid)
	return nil
}

func (service OTPService) VerfiyOTP(to, code string) (bool, error) {
	LoadEnv()

	var client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		return false, err
	}
	if *resp.Status == "approved" {
		return true, nil
	}
	return false, nil
}
