package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers/service_helper"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (handler *userService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.Response, error) {
	//check address is present or not
	res, err := handler.storage.CheckDataExist("addresses", "id", req.GetAddressId())
	if err != nil {
		handler.log.LogError("Error ", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		handler.log.LogError("Error Address not found")
		return nil, status.Errorf(codes.NotFound, utils.NotFound)
	}

	//check cart is present or not
	res, err = handler.storage.CheckDataExist("carts", "id", req.GetCartId())
	if err != nil {
		handler.log.LogError("Error ", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		handler.log.LogError("Error Cart not found")
		return nil, status.Errorf(codes.NotFound, utils.NotFound)
	}

	if req.PromoId != nil {
		//check promoCode is present or not
		res, err = handler.storage.CheckDataExist("promo_codes", "id", req.GetCartId())
		if err != nil {
			handler.log.LogError("Error ", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
		if !res {
			handler.log.LogError("Error Cart not found")
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
	}

	payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve user payload from context")
		handler.log.LogError("Error", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	serviceToken, err := handler.token.CreateServiceToken(payload.UserID, "order")
	if err != nil {
		handler.log.LogError("error while generating service token in CreateOrder", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	conn, err := service_helper.ConnectEndpoints(handler.config.ServerAddress.Order, "order", handler.log)
	if err != nil {
		handler.log.LogError("error while connecting order service :", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	defer conn.Close()

	orderClient := proto.NewOrderServiceClient(conn)
	serviceCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	serviceCtx = metadata.NewOutgoingContext(serviceCtx, metadata.New(map[string]string{
		token.ServiceToken: fmt.Sprint(utils.AuthorizationTypeBearer, " ", serviceToken),
	}))
	defer cancel()

	response, err := orderClient.CreateOrder(serviceCtx, req)
	if err != nil {
		handler.log.LogError("error parsing order service context :", err)
		return nil, err
	}
	return response, nil
}
