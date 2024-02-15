package routes

import (
	"github.com/akmal4410/gestapo/controllers"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/twilio"
	"github.com/akmal4410/gestapo/util"
	"github.com/gorilla/mux"
)

type AuthRoute struct {
	auth *controllers.AuthController
}

var authRoute AuthRoute

func AuthRoutes(router *mux.Router, config *util.Config) {
	twilio := twilio.NewOTPService()
	email := mail.NewGmailService(config.SenderName, config.SenderEmailAddress, config.SemderEmailPassword)

	authRoute.auth = controllers.NewAuthController(twilio, email)

	router.HandleFunc("/send_otp", authRoute.auth.SendOTP).Methods("POST")
	router.HandleFunc("/login", authRoute.auth.LoginUser)
	router.HandleFunc("/create", authRoute.auth.CreateUser)
}
