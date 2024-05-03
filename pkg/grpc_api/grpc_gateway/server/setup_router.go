package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/grpc_api/grpc_gateway/server/middleware"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (server *RestServer) SetupRouter(mux *http.ServeMux) {

	//EditProfile
	editProfile := middleware.ApplyAccessRoleMiddleware(server.token, server.log, utils.MERCHANT, http.HandlerFunc(server.EditProfile))
	mux.Handle("/profile", editProfile) //.Methods("PATCH")

	//InsertProduct
	addProduct := middleware.ApplyAccessRoleMiddleware(server.token, server.log, utils.MERCHANT, http.HandlerFunc(server.InsertProduct))
	mux.Handle("/product", addProduct) //.Methods("POST")

	//EditProduct
	editProduct := middleware.ApplyAccessRoleMiddleware(server.token, server.log, utils.MERCHANT, http.HandlerFunc(server.EditProduct))
	mux.Handle("/product/{id}", editProduct) //.Methods("PATCH")
}
