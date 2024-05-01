package grpc

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/interceptor"
	"github.com/akmal4410/gestapo/pkg/grpc_api/merchant_service/service"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCService(ctx context.Context, storage *database.Storage, config *config.Config, log logger.Logger) error {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.LogFatal("Error while Initializing NewJWTMaker %w", err)
	}
	service := service.NewMerchantService(storage, config, log, tokenMaker)
	authInterceptor := interceptor.NewAuthInterceptor(tokenMaker, log)

	lis, err := net.Listen("tcp", config.ServerAddress.Merchant)
	if err != nil {
		log.LogError("error in listening to port", config.ServerAddress.Merchant, "error:", err)
		return err
	}

	//For handling gRPC and HTTTp
	m := cmux.New(lis)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(ctx, grpcListener, authInterceptor, service) })
	g.Go(func() error { return httpServe(httpListener, service) })
	g.Go(func() error { return m.Serve() })

	return nil
}

func grpcServe(ctx context.Context, lis net.Listener, authInterceptor *interceptor.AuthInterceptor, service *service.MerchantService) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authInterceptor.AuthMiddleware(),
		),
	)

	proto.RegisterMerchantServiceServer(grpcServer, service)
	service.Log.LogInfo("Registreing for reflection")
	reflection.Register(grpcServer)

	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range c {
			service.Log.LogInfo("shutting down grpc server....")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()
	service.Log.LogInfo("Start gRPC server at ", lis.Addr().String())

	return grpcServer.Serve(lis)
}

func httpServe(l net.Listener, service *service.MerchantService) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant", service.EditProfile)
	s := &http.Server{Handler: mux}
	return s.Serve(l)
}

// Check this one out for grpc and http request
// https://drgarcia1986.medium.com/listen-grpc-and-http-requests-on-the-same-port-263c40cb45ff
