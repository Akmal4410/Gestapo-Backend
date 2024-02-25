package server

import (
	"log"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/auth"
	"github.com/akmal4410/gestapo/pkg/api/auth/database"
	"github.com/akmal4410/gestapo/pkg/service/cache"
	"github.com/akmal4410/gestapo/pkg/service/mail"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/service/twilio"
)

type AuthRoute struct {
	auth *auth.AuthHandler
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

	authStore := database.NewAuthStore(server.storage)

	authRoute.auth = auth.NewAuthHandler(twilio, email, authStore, tokenMaker, redis, server.log, server.config)

	// server.router.HandleFunc("/", authRoute.auth.Home)
	authRoutes := server.router.PathPrefix("/auth").Subrouter()

	authRoutes.HandleFunc("/login", authRoute.auth.LoginUser).Methods("POST")
	authRoutes.HandleFunc("/send-otp", authRoute.auth.SendOTP).Methods("POST")
	// authRoutes.HandleFunc("/sso-callback", authRoute.auth.SSOCallback)
	authRoutes.Handle("/sso-auth", auth.SsoMiddleware(server.log, http.HandlerFunc(authRoute.auth.SSOAuth))).Methods("POST")
	authRoutes.Handle("/signup", auth.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(authRoute.auth.SignUpUser))).Methods("POST")
	authRoutes.Handle("/forgot-password", auth.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(authRoute.auth.ForgotPassword))).Methods("POST")
}
