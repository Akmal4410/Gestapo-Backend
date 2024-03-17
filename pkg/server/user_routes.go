package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/user"
	"github.com/akmal4410/gestapo/pkg/api/user/database"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	"github.com/akmal4410/gestapo/pkg/service/token"
)

func (server *Server) userRoutes() {

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}
	userRoute := server.router.PathPrefix("/user").Subrouter()

	store := database.NewUserStore(server.storage)
	handler := user.NewUserHandler(store)

	//GetHome
	userRoute.Handle("/home", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetHome))).Methods("GET")
}
