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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (handler *merchantService) GetProfile(ctx context.Context, req *proto.GetMerchantProfileRequest) (*proto.GetMerchantProfileResponse, error) {
	if req.GetUserId() == "" {
		handler.log.LogError("Error while Getting user id")
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}
	res, err := handler.storage.CheckDataExist("user_data", "id", req.GetUserId())
	if err != nil {
		handler.log.LogError("Error while CheckUserExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		err = fmt.Errorf("account does'nt exist using %s", req.GetUserId())
		handler.log.LogError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	userData, err := handler.storage.GetProfile(req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProfile", err)
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		handler.log.LogError("Error while GetProfile", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if userData.ProfileImage != nil && *userData.ProfileImage != "" {
		url, err := handler.s3.GetPreSignedURL(*userData.ProfileImage)
		if err != nil {
			handler.log.LogError("Error while GetPreSignedURL", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
		userData.ProfileImage = &url
	}

	var dob *timestamppb.Timestamp
	if userData.DOB != nil {
		dob = timestamppb.New(*userData.DOB)
	}

	merchantData := &proto.MerchantResponse{
		Id:           userData.ID,
		ProfileImage: userData.ProfileImage,
		FullName:     userData.FullName,
		UserName:     userData.UserName,
		Phone:        userData.Phone,
		Email:        userData.Email,
		Dob:          dob,
		Gender:       userData.Gender,
		UserType:     userData.UserType,
	}
	respone := &proto.GetMerchantProfileResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Profile fetched sucessfull",
		Data:    merchantData,
	}
	return respone, nil
}

func (handler *merchantService) GetProducts(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductsResponse, error) {
	payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve merchant payload from context")
		handler.log.LogError("Error", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	serviceToken, err := handler.token.CreateServiceToken(payload.UserID, "product")
	if err != nil {
		handler.log.LogError("error while generating service token in GetProducts", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	conn, err := service_helper.ConnectEndpoints(handler.config.ServerAddress.Product, "product", handler.log)
	if err != nil {
		handler.log.LogError("error while connecting product service :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	defer conn.Close()

	productClient := proto.NewProductServiceClient(conn)
	serviceCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	serviceCtx = metadata.NewOutgoingContext(serviceCtx, metadata.New(map[string]string{
		token.ServiceToken: fmt.Sprint(utils.AuthorizationTypeBearer, " ", serviceToken),
	}))
	defer cancel()

	response, err := productClient.GetProducts(serviceCtx, &proto.GetProductRequest{MerchantId: req.MerchantId})
	if err != nil {
		handler.log.LogError("error parsing product service context :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	return response, nil
}

func (handler *merchantService) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.Response, error) {
	payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve merchant payload from context")
		handler.log.LogError("Error", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	serviceToken, err := handler.token.CreateServiceToken(payload.UserID, "product")
	if err != nil {
		handler.log.LogError("error while generating service token in DeleteProduct", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	conn, err := service_helper.ConnectEndpoints(handler.config.ServerAddress.Product, "product", handler.log)
	if err != nil {
		handler.log.LogError("error while connecting product service :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	defer conn.Close()

	productClient := proto.NewProductServiceClient(conn)
	serviceCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	serviceCtx = metadata.NewOutgoingContext(serviceCtx, metadata.New(map[string]string{
		token.ServiceToken: fmt.Sprint(utils.AuthorizationTypeBearer, " ", serviceToken),
	}))
	defer cancel()

	productRes, err := productClient.GetProductById(serviceCtx, &proto.GetProductByIdRequest{
		ProductId: req.GetProductId(),
	})
	if err != nil {
		handler.log.LogError("Error while retrieving product", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !productRes.Status {
		if productRes.Code == int32(codes.NotFound) {
			handler.log.LogError("Error while GetProductById product Not found")
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		handler.log.LogError("Error while GetProductById")
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if productRes.Data.MerchantId != nil && *productRes.Data.MerchantId != payload.UserID {
		handler.log.LogError("unauthorized: product does not belong to the authenticated merchant")
		return nil, status.Errorf(codes.PermissionDenied, "product does not belong to the authenticated merchant")
	}

	err = handler.storage.DeleteProduct(req.GetProductId())
	if err != nil {
		handler.log.LogError("Error while DeleteProduct", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	for _, key := range productRes.Data.ProductImages {
		err := handler.s3.DeleteKey(key)
		if err != nil {
			handler.log.LogError("Error deleting file from S3", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
	}

	response := &proto.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Product deleted succesfully",
	}
	return response, nil
}

func (handler *merchantService) AddProductDiscount(ctx context.Context, in *proto.AddDiscountRequest) (*proto.Response, error) {
	payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve merchant payload from context")
		handler.log.LogError("Error", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	req := &entity.AddDiscountReq{
		ProductId:    in.GetProductId(),
		MerchantId:   payload.UserID,
		DiscountName: in.GetName(),
		Description:  in.GetDescription(),
		Percentage:   float64(in.GetPercentage()),
		CardColor:    in.GetCardColor(),
		StartTime:    in.GetStartTime().AsTime(),
		EndTime:      in.GetEndTime().AsTime(),
	}
	err := helpers.ValidateBody(nil, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}

	serviceToken, err := handler.token.CreateServiceToken(payload.UserID, "product")
	if err != nil {
		handler.log.LogError("error while generating service token in DeleteProduct", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	conn, err := service_helper.ConnectEndpoints(handler.config.ServerAddress.Product, "product", handler.log)
	if err != nil {
		handler.log.LogError("error while connecting product service :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	defer conn.Close()

	productClient := proto.NewProductServiceClient(conn)
	serviceCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	serviceCtx = metadata.NewOutgoingContext(serviceCtx, metadata.New(map[string]string{
		token.ServiceToken: fmt.Sprint(utils.AuthorizationTypeBearer, " ", serviceToken),
	}))
	defer cancel()

	productRes, err := productClient.GetProductById(serviceCtx, &proto.GetProductByIdRequest{
		ProductId: in.GetProductId(),
	})
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = No found" {
			handler.log.LogError("Error while GetProductById Not found", err)
			return nil, status.Errorf(codes.NotFound, "No found")
		}
		handler.log.LogError("Error while retrieving product", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !productRes.Status {
		if productRes.Code == int32(codes.NotFound) {
			handler.log.LogError("Error while GetProductById product Not found")
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		handler.log.LogError("Error while GetProductById")
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if productRes.Data.MerchantId != nil && *productRes.Data.MerchantId != payload.UserID {
		handler.log.LogError("unauthorized: product does not belong to the authenticated merchant")
		return nil, status.Errorf(codes.PermissionDenied, "product does not belong to the authenticated merchant")
	}

	err = handler.storage.AddProductDiscount(req)
	if err != nil {
		handler.log.LogError("Error while AddProductDiscount", err)
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

func (handler *merchantService) GetAllDiscounts(ctx context.Context, req *proto.GetDiscountsRequest) (*proto.GetDiscountsResponse, error) {
	discountEntities, err := handler.storage.GetAllDiscount(req.MerchantId)
	if err != nil {
		handler.log.LogError("Error while GetAllDiscount", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	var discouts []*proto.DiscountResponse
	for _, discount := range discountEntities {
		discountRes := &proto.DiscountResponse{
			ProductId:    discount.ProductID,
			Name:         discount.Name,
			Description:  discount.Description,
			Percentage:   float32(discount.Percentage),
			ProductImage: discount.ProductImage,
			CardColor:    discount.CardColor,
		}
		discouts = append(discouts, discountRes)
	}
	response := &proto.GetDiscountsResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Discounts fetched successfully",
		Data:    discouts,
	}
	return response, nil
}
