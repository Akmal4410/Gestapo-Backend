package grpc

import (
	"net"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/service"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCService(storage *database.Storage, config *config.Config, log logger.Logger) error {

	grpcServer := grpc.NewServer()
	service := service.NewAuthenticationService(storage, config, log)

	proto.RegisterAuthenticationServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", config.ServerAddress.Authentication)
	if err != nil {
		log.LogError("error in listening to port", config.ServerAddress.Authentication, "error:", err)
		return err
	}
	log.LogInfo("Start gRPC server at : %s", lis.Addr().String())
	return grpcServer.Serve(lis)
}
