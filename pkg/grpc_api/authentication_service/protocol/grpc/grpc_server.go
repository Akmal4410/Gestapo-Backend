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
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/service"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCService(ctx context.Context, storage *database.Storage, config *config.Config, log logger.Logger) error {

	grpcServer := grpc.NewServer()
	service := service.NewAuthenticationService(storage, config, log)

	proto.RegisterAuthenticationServiceServer(grpcServer, service)
	log.LogInfo("Registreing for reflection")
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", config.ServerAddress.Authentication)
	if err != nil {
		log.LogError("error in listening to port", config.ServerAddress.Authentication, "error:", err)
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
	log.LogInfo("Start gRPC server at : %s", lis.Addr().String())
	return grpcServer.Serve(lis)
}
