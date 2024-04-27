package auth_interceptor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	signUp         string = "/pb.AuthenticationService/SignUpUser"
	forgotPassword string = "auth/forgot-password"
)

type AuthInterceptor struct {
	token token.Maker
	log   logger.Logger
}

func NewAuthInterceptor(token token.Maker, log logger.Logger) *AuthInterceptor {
	return &AuthInterceptor{
		token: token,
		log:   log,
	}
}

// AuthServerInterceptor is a gRPC unary server interceptor for authentication.
func (interceptor *AuthInterceptor) AuthServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		interceptor.log.LogInfo("Calling gRPC meathod :", info.FullMethod)
		if ok := IsAuthenticationNeeded(info.FullMethod); !ok {
			return handler(ctx, req)
		}
		response := &proto.Response{}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err := errors.New("metadata is not provided")
			interceptor.log.LogError("Error :", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusUnauthorized, err.Error())
			return response, nil
		}

		authorizationHeaders := md.Get(utils.AuthorizationKey)
		if len(authorizationHeaders) == 0 {
			err := errors.New("authorization header is not provided")
			interceptor.log.LogError("Error", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusUnauthorized, err.Error())
			return response, nil
		}

		authorizationHeader := authorizationHeaders[0]
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			interceptor.log.LogError("Error", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusUnauthorized, err.Error())
			return response, nil
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authorizationType)
			interceptor.log.LogError("Error", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusUnauthorized, err.Error())
			return response, nil
		}

		token := fields[1]
		// Verify and parse the token
		payload, err := interceptor.token.VerifySessionToken(token)
		if err != nil {
			err := fmt.Errorf("error while VerifySessionToken: %s", err.Error())
			interceptor.log.LogError("Error", err)
			response.ErrorInfo = helpers.ErrorJson(http.StatusUnauthorized, err.Error())
			return response, nil
		}

		// Add the payload to the context
		ctx = context.WithValue(ctx, utils.AuthorizationPayloadKey, payload)

		return handler(ctx, req)
	}
}

// IsAuthenticationNeeded returns true if the route needed authentication middleware
func IsAuthenticationNeeded(route string) bool {
	switch route {
	case signUp, forgotPassword:
		return true
	}
	return false
}
