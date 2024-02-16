package routes

import (
	"log"

	"github.com/akmal4410/gestapo/controllers"
	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/token"
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
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("%w", err)
	}
	authRoute.auth = controllers.NewAuthController(twilio, email, storage, tokenMaker)

	router.HandleFunc("/send_otp", authRoute.auth.SendOTP).Methods("POST")
	router.HandleFunc("/login", authRoute.auth.LoginUser)
	router.HandleFunc("/create", authRoute.auth.CreateUser)
}
