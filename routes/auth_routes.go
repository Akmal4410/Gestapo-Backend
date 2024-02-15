package routes

import (
	"github.com/akmal4410/gestapo/controllers"
	"github.com/gorilla/mux"
)

type AuthRoute struct {
	login controllers.Auth
}

var authRoute AuthRoute

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/send_otp", authRoute.login.SendOTP).Methods("POST")
	router.HandleFunc("/login", authRoute.login.LoginUser)
	router.HandleFunc("/create", authRoute.login.CreateUser)
}
