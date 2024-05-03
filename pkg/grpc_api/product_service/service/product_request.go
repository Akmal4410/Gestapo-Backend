package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers/service_helper"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (handler *productService) GetProducts(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductsResponse, error) {
	_, err := service_helper.ValidateServiceToken(ctx, handler.log, handler.token)
	if err != nil {
		handler.log.LogError("Error while ValidateServiceToken", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	productRes, err := handler.storage.GetProducts()
	if err != nil {
		handler.log.LogError("Error while GetProducts", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	for _, product := range productRes {
		if product.ProductImages != nil {
			for i, image := range product.ProductImages {
				url, err := handler.s3.GetPreSignedURL(image)
				if err != nil {
					handler.log.LogError("Error while GetPreSignedURL", err)
					return nil, status.Errorf(codes.Internal, utils.InternalServerError)
				}
				product.ProductImages[i] = url
			}
		}
	}
	var products []*proto.ProductResponse
	for _, product := range productRes {
		newProduct := &proto.ProductResponse{
			Id:            product.ID,
			MerchantId:    product.MerchantID,
			ProductImages: product.ProductImages,
			ProductName:   *product.ProductName,
			Description:   *product.Description,
			CategoryName:  *product.CategoryName,
			Size:          *product.Size,
			Price:         product.Price,
			DiscountPrice: *product.DiscountPrice,
		}
		products = append(products, newProduct)
	}

	response := &proto.GetProductsResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Products fetched successfully",
		Data:    products,
	}
	return response, nil
}

func (handler *productService) GetProductById(ctx context.Context, req *proto.GetProductByIdRequest) (*proto.GetProductByIdResponse, error) {
	product, err := handler.storage.GetProductById(req.GetProductId())
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProductById Not found", err)
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		handler.log.LogError("Error while GetProductById", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if product.ProductImages != nil {
		for i, image := range product.ProductImages {
			url, err := handler.s3.GetPreSignedURL(image)
			if err != nil {
				handler.log.LogError("Error while GetPreSignedURL", err)
				return nil, status.Errorf(codes.Internal, utils.InternalServerError)
			}
			product.ProductImages[i] = url
		}
	}

	productRes := &proto.ProductResponse{
		Id:            product.ID,
		MerchantId:    product.MerchantID,
		ProductImages: product.ProductImages,
		ProductName:   *product.ProductName,
		Description:   *product.Description,
		CategoryName:  *product.CategoryName,
		Size:          *product.Size,
		Price:         product.Price,
		DiscountPrice: *product.DiscountPrice,
	}

	response := &proto.GetProductByIdResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Product fetched successfully",
		Data:    productRes,
	}

	return response, nil
}