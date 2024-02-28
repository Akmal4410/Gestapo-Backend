package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/merchant"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	"github.com/akmal4410/gestapo/pkg/service/token"
)

func (server *Server) merchantRoutes() {

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}

	store := database.NewMarchantStore(server.storage)

	handler := merchant.NewMerchentHandler(store, server.log)

	merchantRoutes := server.router.PathPrefix("/merchant").Subrouter()

	merchantRoutes.Handle("/profile/{id}", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProfile))).Methods("GET")
	merchantRoutes.Handle("/profile", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.EditProfile))).Methods("PATCH")
}
