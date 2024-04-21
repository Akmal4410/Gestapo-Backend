package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/user"
	"github.com/akmal4410/gestapo/pkg/api/user/database"
	db "github.com/akmal4410/gestapo/pkg/database"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
)

func (server *Server) userRoutes() {

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}
	userRoute := server.router.PathPrefix("/user").Subrouter()

	store := database.NewUserStore(server.storage)
	dbStore := db.NewDBStore(server.storage)
	handler := user.NewUserHandler(server.log, store, dbStore, server.s3)

	//GetHome
	userRoute.Handle("/home", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetHome))).Methods("GET")
}
