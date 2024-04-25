package grpc_gateway

import (
	"context"
	"net"
	"net/http"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
)

const (
	serviceName = "gRPC Gateway"
	logFileName = "grpc_gateway"
)

func RunGateway() error {
	log := logger.NewLogrusLogger(logFileName)
	log.LogInfo(serviceName, "has started")

	config, err := config.LoadConfig("configs")
	if err != nil {
		log.LogFatal("Cannot load configuration:", err)
	}
	log.LogInfo("Config file loaded.")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gMux, err := newGateway(ctx, log, config)
	if err != nil {
		log.LogError("error in newGateway :", err)
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", gMux))

	lis, err := net.Listen("tcp", config.ServerAddress.Gateway)
	if err != nil {
		log.LogError("error in listening to port", config.ServerAddress.Gateway, "error:", err)
		return err
	}
	log.LogInfo("Listening to port", config.ServerAddress.Gateway)
	err = http.Serve(lis, mux)
	if err != nil {
		log.LogError("Cannot server gateway", config.ServerAddress.Gateway, "error:", err)
		return err
	}
	return nil

}
