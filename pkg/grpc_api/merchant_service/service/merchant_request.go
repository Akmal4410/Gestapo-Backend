package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/merchant_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/service_helper"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (merchant *merchantService) GetProfile(ctx context.Context, req *proto.GetMerchantProfileRequest) (*proto.GetMerchantProfileResponse, error) {
	if req.GetUserId() == "" {
		merchant.log.LogError("Error while Getting user id")
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}
	res, err := merchant.storage.CheckDataExist("user_data", "id", req.GetUserId())
	if err != nil {
		merchant.log.LogError("Error while CheckUserExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		err = fmt.Errorf("account does'nt exist using %s", req.GetUserId())
		merchant.log.LogError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	userData, err := merchant.storage.GetProfile(req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			merchant.log.LogError("Error while GetProfile", err)
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		merchant.log.LogError("Error while GetProfile", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if userData.ProfileImage != nil {
		url, err := merchant.s3.GetPreSignedURL(*userData.ProfileImage)
		if err != nil {
			merchant.log.LogError("Error while GetPreSignedURL", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
		userData.ProfileImage = &url
	}
	respone := &proto.GetMerchantProfileResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Profile fetched sucessfull",
		Data:    userData,
	}
	return respone, nil
}

func (merchant *merchantService) GetProducts(ctx context.Context, req *proto.Request) (*proto.GetProductsResponse, error) {
	payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve merchant payload from context")
		merchant.log.LogError("Error", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	serviceToken, err := merchant.token.CreateServiceToken(payload.UserID, "authentication")
	if err != nil {
		merchant.log.LogError("error while generating service token in GetProducts", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	conn, err := service_helper.ConnectEndpoints(merchant.config.ServerAddress.Product, "product", merchant.log)
	if err != nil {
		merchant.log.LogError("error while connecting product service :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	defer conn.Close()

	productClient := proto.NewProductServiceClient(conn)
	serviceCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	serviceCtx = metadata.NewOutgoingContext(serviceCtx, metadata.New(map[string]string{
		"service-token": serviceToken,
	}))
	defer cancel()

	response, err := productClient.GetProducts(serviceCtx, &proto.GetProductRequest{MerchantId: payload.UserID})
	if err != nil {
		merchant.log.LogError("error parsing product service context : :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	return response, nil

	// res, err := merchant.dbStorage.GetProducts()
	// if err != nil {
	// 	merchant.log.LogError("Error while GetProducts", err)
	// 	return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	// }
	// for _, product := range res {
	// 	if product.ProductImages != nil {
	// 		for i, image := range product.ProductImages {
	// 			url, err := merchant.s3.GetPreSignedURL(image)
	// 			if err != nil {
	// 				merchant.log.LogError("Error while GetPreSignedURL", err)
	// 				return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	// 			}
	// 			product.ProductImages[i] = url
	// 		}
	// 	}
	// }

	// helpers.WriteJSON(w, http.StatusOK, res)
}

//FOR PRODUCT Service

// func (handler *MerchantHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]

// 	product, err := handler.dbStorage.GetProductById(productId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProductById Not found", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while GetProductById", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if product.ProductImages != nil {
// 		for i, image := range product.ProductImages {
// 			url, err := handler.s3Service.GetPreSignedURL(image)
// 			if err != nil {
// 				handler.log.LogError("Error while GetPreSignedURL", err)
// 				helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 				return
// 			}
// 			product.ProductImages[i] = url
// 		}
// 	}
// 	helpers.WriteJSON(w, http.StatusOK, product)
// }

//FOR PRODUCT Service

// func (handler *MerchantHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	product, err := handler.dbStorage.GetProductById(productId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProductById Not fount", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while retrieving product", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	payload, ok := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
// 	if !ok {
// 		err := errors.New("unable to retrieve user payload from context")
// 		handler.log.LogError("Error", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if product.MerchantID != payload.UserID {
// 		err := errors.New("unauthorized: product does not belong to the authenticated merchant")
// 		handler.log.LogError("Error", err)
// 		helpers.ErrorJson(http.StatusForbidden, err.Error())
// 		return
// 	}

// 	err = handler.storage.DeleteProduct(productId)
// 	if err != nil {
// 		handler.log.LogError("Error while DeleteProduct", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	for _, key := range product.ProductImages {
// 		err := handler.s3Service.DeleteKey(key)
// 		if err != nil {
// 			handler.log.LogError("Error deleting file from S3", err)
// 			helpers.ErrorJson(http.StatusInternalServerError, "Error deleting file from")
// 			return
// 		}
// 	}
// 	helpers.WriteJSON(w, http.StatusOK, "Product deleted succesfully")
// }

func (merchant *merchantService) AddProductDiscount(ctx context.Context, in *proto.AddDiscountRequest) (*proto.Response, error) {
	req := &entity.AddDiscountReq{
		ProductId:    in.GetProductId(),
		DiscountName: in.GetName(),
		Description:  in.GetDescription(),
		Percentage:   float64(in.GetPercentage()),
		CardColor:    in.GetCardColor(),
		StartTime:    in.GetStartTime().AsTime(),
		EndTime:      in.GetEndTime().AsTime(),
	}
	err := helpers.ValidateBody(nil, req)
	if err != nil {
		merchant.log.LogError("Error while ValidateBody", err)
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}
	exist, err := merchant.storage.CheckDataExist("products", "id", req.ProductId)
	if err != nil {
		merchant.log.LogError("Error while checking product", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !exist {
		merchant.log.LogError("product doesnt exists")
		return nil, status.Errorf(codes.NotFound, "Product not found with ", req.ProductId)
	}

	err = merchant.storage.AddProductDiscount(req)
	if err != nil {
		merchant.log.LogError("Error while AddProductDiscount", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)

	}
	response := &proto.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Discount added successfully",
	}

	return response, nil
}

func (handler *merchantService) EditProductDiscount(ctx context.Context, in *proto.EditDiscountRequest) (*proto.Response, error) {
	startTime := in.GetStartTime().AsTime()
	entTime := in.GetEndTime().AsTime()
	req := &entity.EditDiscountReq{
		DiscountName: &in.Name,
		Description:  &in.Description,
		Percentage:   float64(in.GetPercentage()),
		CardColor:    &in.CardColor,
		StartTime:    &startTime,
		EndTime:      &entTime,
	}
	err := helpers.ValidateBody(nil, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}

	res, err := handler.storage.CheckDataExist("discounts", "id", in.GetDiscountId())
	if err != nil {
		handler.log.LogError("Error while CheckDataExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		err = fmt.Errorf("discounts doesnt exist: %s", in.GetDiscountId())
		handler.log.LogError("Error ", err)
		return nil, status.Errorf(codes.NotFound, utils.NotFound)
	}

	err = handler.storage.EditProductDiscount(in.GetDiscountId(), req)
	if err != nil {
		handler.log.LogError("Error while ApplyProductDiscount", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	response := &proto.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Discount updated successfully",
	}

	return response, nil
}
