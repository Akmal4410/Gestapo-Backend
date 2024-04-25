package grpc_gateway

import (
	"context"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newGateway(ctx context.Context, log logger.Logger, config config.Config, opts ...runtime.ServeMuxOption) (*runtime.ServeMux, error) {
	gMux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	errAuthentication := registerAuthenticationEndPoints(ctx, log, config, gMux, dialOpts)
	if errAuthentication != nil {
		log.LogError("error in registering Authentication Service.", errAuthentication)
		return nil, errAuthentication
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
