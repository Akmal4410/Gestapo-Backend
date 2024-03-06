package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/admin"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (server *Server) adminRoutes() {
	adminHandler := admin.NewAdminHandler(server.log)

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}

	adminRoutes := server.router.PathPrefix("/admin").Subrouter()
	category := middleware.ApplyMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.CreateCategories))
	adminRoutes.Handle("/category", category).Methods("POST")
}
