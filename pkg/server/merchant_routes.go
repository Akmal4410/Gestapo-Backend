package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/merchant"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	db "github.com/akmal4410/gestapo/pkg/database"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (server *Server) merchantRoutes() {

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}

	store := database.NewMarchantStore(server.storage)
	dbStore := db.NewDBStore(server.storage)
	handler := merchant.NewMerchentHandler(server.log, store, dbStore, server.s3)

	merchantRoutes := server.router.PathPrefix("/merchant").Subrouter()

	//GetProfile
	merchantRoutes.Handle("/profile/{id}", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProfile))).Methods("GET")

	//EditProfile
	editProfile := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.EditProfile))
	merchantRoutes.Handle("/profile", editProfile).Methods("PATCH")

	//InsertProduct
	addProduct := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.InsertProduct))
	merchantRoutes.Handle("/product", addProduct).Methods("POST")

	//EditProduct
	editProduct := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.EditProduct))
	merchantRoutes.Handle("/product/{id}", editProduct).Methods("PATCH")

	//GetProducts
	merchantRoutes.Handle("/product", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProducts))).Methods("GET")

	//GetProductById
	merchantRoutes.Handle("/product/{id}", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProductById))).Methods("GET")

	//DeleteProduct
	deleteProduct := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.DeleteProduct))
	merchantRoutes.Handle("/product/{id}", deleteProduct).Methods("DELETE")

	//AddProductDiscount
	addProductDiscount := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.AddProductDiscount))
	merchantRoutes.Handle("/product/discount", addProductDiscount).Methods("POST")

	editProductDiscount := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.EditProductDiscount))
	merchantRoutes.Handle("/product/discount/{id}", editProductDiscount).Methods("PATCH")
}
