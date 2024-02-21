package routes

import (
	"log"
	"net/http"

	"github.com/akmal4410/gestapo/controllers"
	"github.com/akmal4410/gestapo/middlewares"
	"github.com/akmal4410/gestapo/services/cache"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/token"
	"github.com/akmal4410/gestapo/services/twilio"
)

type AuthRoute struct {
	auth *controllers.AuthController
}

var authRoute AuthRoute

func (server *Server) authRoutes() {
	twilio := twilio.NewOTPService()
	email := mail.NewGmailService(server.config.Email)
	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("%w", err)
	}
	redis, err := cache.NewRedisCache(server.config.Redis)
	if err != nil {
		log.Fatal("%w", err)
	}

	authRoute.auth = controllers.NewAuthController(twilio, email, server.storage, tokenMaker, redis, server.log)

	authRoutes := server.router.PathPrefix("/auth").Subrouter()

	authRoutes.HandleFunc("/login", authRoute.auth.LoginUser).Methods("POST")
	authRoutes.HandleFunc("/send-otp", authRoute.auth.SendOTP).Methods("POST")
	authRoutes.HandleFunc("/verify-otp", authRoute.auth.VerifyOTP).Methods("POST")
	authRoutes.HandleFunc("/forgot-password", authRoute.auth.LoginUser).Methods("POST")
	authRoutes.Handle("/signup", middlewares.AuthenticationMiddleware(tokenMaker, server.log, http.HandlerFunc(authRoute.auth.SignUpUser))).Methods("POST")
}
