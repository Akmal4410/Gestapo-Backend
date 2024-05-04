package interceptor

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	getProductRPC string = "/pb.ProductService/GetProducts"
)

// For protecting merchant calls
const (
	deletProduct        string = "/pb.MerchantService/DeleteProduct"
	addProductDiscount  string = "/pb.MerchantService/AddProductDiscount"
	editProductDiscount string = "/pb.MerchantService/EditProductDiscount"
)

// AccessMiddleware is a gRPC unary server interceptor for access.
func (interceptor *Interceptor) AccessMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		interceptor.log.LogInfo("Calling gRPC meathod :", info.FullMethod)
		//	skipping the authentication middlware because it don't have access token
		//	so add method names that wanted to skip inside this
		if ok := skipAuthenticationBetweenRPC(info.FullMethod); ok {
			return handler(ctx, req)
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err := errors.New("metadata is not provided")
			interceptor.log.LogError("Error : ", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		authorizationHeaders := md.Get(utils.AuthorizationKey)
		if len(authorizationHeaders) == 0 {
			err := errors.New("authorization header is not provided")
			interceptor.log.LogError("Error : ", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		authorizationHeader := authorizationHeaders[0]
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			interceptor.log.LogError("Error : ", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authorizationType)
			interceptor.log.LogError("Error : ", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		token := fields[1]
		// Verify and parse the token
		payload, err := interceptor.token.VerifyAccessToken(token)
		if err != nil {
			err := fmt.Errorf("error while VerifyAccessToken: %s", err.Error())
			interceptor.log.LogError("Error : ", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		if payload.TokenType != "access-token" {
			err := fmt.Errorf("invalid token type: %s", payload.TokenType)
			interceptor.log.LogError("Error", err)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())

		}

		// Add the payload to the context
		ctx = context.WithValue(ctx, utils.AuthorizationPayloadKey, payload)

		return handler(ctx, req)
	}
}

// Currently used by merchant so modify if needed
func (interceptor *Interceptor) RolMiddleware(requiredRole string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		interceptor.log.LogInfo("Calling gRPC meathod :", info.FullMethod)
		if ok := isMerchantCanOnlyAccess(info.FullMethod); !ok {
			return handler(ctx, req)
		}
		payload, ok := ctx.Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
		if !ok {
			err := errors.New("unable to retrieve user payload from context")
			interceptor.log.LogError("Error", err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		if payload.UserType != requiredRole {
			err := fmt.Errorf("user does not have required role: %s", requiredRole)
			interceptor.log.LogError("Error", err)
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}
		ctx = context.WithValue(ctx, utils.AuthorizationPayloadKey, payload)
		return handler(ctx, req)
	}
}

func isMerchantCanOnlyAccess(method string) bool {
	switch method {
	case deletProduct, addProductDiscount, editProductDiscount:
		return true
	}
	return false
}

func skipAuthenticationBetweenRPC(method string) bool {
	switch method {
	case getProductRPC:
		return true
	}
	return false
}
