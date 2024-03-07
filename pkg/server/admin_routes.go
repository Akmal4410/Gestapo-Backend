package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/admin"
	"github.com/akmal4410/gestapo/pkg/api/admin/database"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (server *Server) adminRoutes() {
	storage := database.NewAdminStore(server.storage)
	adminHandler := admin.NewAdminHandler(storage, server.log)

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}

	adminRoutes := server.router.PathPrefix("/admin").Subrouter()

	createCategory := middleware.ApplyMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.CreateCategory))
	adminRoutes.Handle("/category", createCategory).Methods("POST")
	getCategory := middleware.ApplyMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.GetCategories))
	adminRoutes.Handle("/category", getCategory).Methods("GET")

	allUsers := middleware.ApplyMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.GetUsers))
	adminRoutes.Handle("/user", allUsers).Methods("GET")
}
