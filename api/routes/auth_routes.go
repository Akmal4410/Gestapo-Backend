package routes

import (
	"log"
	"net/http"

	"github.com/akmal4410/gestapo/api/handler"
	"github.com/akmal4410/gestapo/api/middleware"
	"github.com/akmal4410/gestapo/pkg/service/cache"
	"github.com/akmal4410/gestapo/pkg/service/mail"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/service/twilio"
)

type AuthRoute struct {
	auth *handler.AuthHandler
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

	authRoute.auth = handler.NewAuthHandler(twilio, email, server.storage, tokenMaker, redis, server.log, server.config)

	// server.router.HandleFunc("/", authRoute.auth.Home)
	authRoutes := server.router.PathPrefix("/auth").Subrouter()

	authRoutes.HandleFunc("/login", authRoute.auth.LoginUser).Methods("POST")
	authRoutes.HandleFunc("/send-otp", authRoute.auth.SendOTP).Methods("POST")
	// authRoutes.HandleFunc("/sso-callback", authRoute.auth.SSOCallback)
	authRoutes.Handle("/sso-auth", middleware.SsoMiddleware(server.log, http.HandlerFunc(authRoute.auth.SSOAuth))).Methods("POST")
	authRoutes.Handle("/signup", middleware.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(authRoute.auth.SignUpUser))).Methods("POST")
	authRoutes.Handle("/forgot-password", middleware.AuthMiddleware(tokenMaker, server.log, http.HandlerFunc(authRoute.auth.ForgotPassword))).Methods("POST")

}
