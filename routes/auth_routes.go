package routes

import (
	"github.com/akmal4410/gestapo/controllers"
	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/twilio"
	"github.com/akmal4410/gestapo/utils"
	"github.com/gorilla/mux"
)

type AuthRoute struct {
	auth *controllers.AuthController
}

var authRoute AuthRoute

func AuthRoutes(router *mux.Router, config *utils.Config, storage *database.Storage) {
	twilio := twilio.NewOTPService()
	email := mail.NewGmailService(config.SenderName, config.SenderEmailAddress, config.SemderEmailPassword)

	authRoute.auth = controllers.NewAuthController(twilio, email, storage)

	router.HandleFunc("/send_otp", authRoute.auth.SendOTP).Methods("POST")
	router.HandleFunc("/login", authRoute.auth.LoginUser)
	router.HandleFunc("/create", authRoute.auth.CreateUser)
}
