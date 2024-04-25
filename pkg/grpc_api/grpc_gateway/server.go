package grpc_gateway

import (
	"context"
	"net/http"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/gorilla/handlers"
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

	// lis, err := net.Listen("tcp", config.ServerAddress.Gateway)
	// if err != nil {
	// 	log.LogError("error in listening to port", config.ServerAddress.Gateway, "error:", err)
	// 	return err
	// }
	// err = http.Serve(lis, mux)
	// if err != nil {
	// 	log.LogError("Cannot server gateway", config.ServerAddress.Gateway, "error:", err)
	// 	return err
	// }

	return http.ListenAndServe(":"+"9000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "User-Agent"}),
			handlers.ExposedHeaders([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(mux),
	)
}
