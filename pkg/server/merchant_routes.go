package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/merchant"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/server/middleware"
	s3service "github.com/akmal4410/gestapo/pkg/service/s3_service"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (server *Server) merchantRoutes() {

	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
	if err != nil {
		server.log.LogFatal("%w", err)
	}
	s3 := s3service.NewS3Service(
		server.config.AwsS3.BucketName,
		server.config.AwsS3.Region,
		server.config.AwsS3.AccessKey,
		server.config.AwsS3.SecretKey,
	)

	store := database.NewMarchantStore(server.storage)

	handler := merchant.NewMerchentHandler(store, server.log, s3)

	merchantRoutes := server.router.PathPrefix("/merchant").Subrouter()

	//GetProfile
	merchantRoutes.Handle("/profile/{id}", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProfile))).Methods("GET")

	//EditProfile
	editProfile := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.EditProfile))
	merchantRoutes.Handle("/profile", editProfile).Methods("PATCH")

	//InsertProduct
	addProduct := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.MERCHANT, http.HandlerFunc(handler.InsertProduct))
	merchantRoutes.Handle("/product", addProduct).Methods("POST")

	//GetProducts
	merchantRoutes.Handle("/product", middleware.AccessMiddleware(tokenMaker, server.log, http.HandlerFunc(handler.GetProducts))).Methods("GET")
}
