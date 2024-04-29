package grpc_gateway

import (
	"context"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func newGateway(ctx context.Context, log logger.Logger, config config.Config, opts ...runtime.ServeMuxOption) (*runtime.ServeMux, error) {
	muxOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	gMux := runtime.NewServeMux(muxOption)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	errAuthentication := registerAuthenticationEndPoints(ctx, log, config, gMux, dialOpts)
	if errAuthentication != nil {
		return nil, errAuthentication
	}

	errAdmin := registerAdminEndPoints(ctx, log, config, gMux, dialOpts)
	if errAuthentication != nil {
		return nil, errAdmin
	}

	return gMux, nil
}

func registerAuthenticationEndPoints(ctx context.Context, log logger.Logger, config config.Config, gMux *runtime.ServeMux, dialOpts []grpc.DialOption) error {
	var endpoint *string
	if config.ServerAddress != nil {
		endpoint = &config.ServerAddress.Authentication
		err := proto.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, gMux, *endpoint, dialOpts)
		if err != nil {
			log.LogError("error in registering authentication endpoint.", err)
			return err
		}
	}
	return nil
}

func registerAdminEndPoints(ctx context.Context, log logger.Logger, config config.Config, gMux *runtime.ServeMux, dialOpts []grpc.DialOption) error {
	var endpoint *string
	if config.ServerAddress != nil {
		endpoint = &config.ServerAddress.Admin
		err := proto.RegisterAdminServiceHandlerFromEndpoint(ctx, gMux, *endpoint, dialOpts)
		if err != nil {
			log.LogError("error in registering admin endpoint.", err)
			return err
		}
	}
	return nil
}
