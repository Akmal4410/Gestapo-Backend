package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/auth"
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/db"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/service/cache"
	"github.com/akmal4410/gestapo/pkg/service/mail"
	"github.com/akmal4410/gestapo/pkg/service/twilio"
)

func (server *Server) authRoutes() {
	twilio := twilio.NewOTPService()
	email := mail.NewGmailService(server.config.Email)
	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}
	redis, err := cache.NewRedisCache(server.config.Redis)
	if err != nil {
		server.log.LogFatal("%w", err)
	}

	authStore := db.NewAuthStore(server.storage)

	handler := auth.NewAuthHandler(twilio, email, authStore, tokenMaker, redis, server.log, server.config)

	// server.router.HandleFunc("/", authRoute.auth.Home)
	authRoutes := server.router.PathPrefix("/auth").Subrouter()

	// authRoutes.HandleFunc("/sso-callback", authRoute.auth.SSOCallback)
	// authRoutes.Handle("/sso-auth", auth.SsoMiddleware(server.log, http.HandlerFunc(handler.SSOAuth))).Methods("POST")
	// authRoutes.Handle("/signup", auth.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.SignUpUser))).Methods("POST")
	authRoutes.Handle("/forgot-password", auth.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.ForgotPassword))).Methods("POST")
}
