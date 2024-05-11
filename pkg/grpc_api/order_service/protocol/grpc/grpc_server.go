package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/order_service/service"
	"github.com/akmal4410/gestapo/pkg/helpers/interceptor"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCService(ctx context.Context, storage *database.Storage, config *config.Config, log logger.Logger) error {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.LogFatal("Error while Initializing NewJWTMaker %w", err)
	}
	service := service.NewOrderService(storage, config, log, tokenMaker)
	interceptor := interceptor.NewInterceptor(tokenMaker, log)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.AccessMiddleware(),
		),
	)

	proto.RegisterOrderServiceServer(grpcServer, service)
	log.LogInfo("Registreing for reflection")
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", config.ServerAddress.Order)
	if err != nil {
		log.LogError("error in listening to port", config.ServerAddress.Order, "error:", err)
		return err
	}
	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range c {
			log.LogInfo("shutting down grpc server....")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()
	log.LogInfo("Start gRPC server at ", lis.Addr().String())
	return grpcServer.Serve(lis)
}
