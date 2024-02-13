package routes

import (
	"github.com/akmal4410/gestapo/controller"
	"github.com/gorilla/mux"
)

type AuthRoute struct {
	login controller.AuthController
}

var authRoute AuthRoute

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", authRoute.login.LoginUser)
	router.HandleFunc("/create", authRoute.login.CreateUser)
}
