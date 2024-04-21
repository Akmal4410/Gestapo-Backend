package authentication_service

import (
	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/protocol/grpc"
	"github.com/akmal4410/gestapo/pkg/helpers/service_helper"
)

const (
	serviceName = "Authentication Service"
	logFileName = "authentication_service"
)

func RunServer() error {
	_, log := service_helper.InitializeService(serviceName, logFileName)
	config, err := config.LoadConfig("configs")
	if err != nil {
		log.LogFatal("Cannot load configuration:", err)
	}
	log.LogInfo("Config file loaded.")

	store, err := database.NewStorage(config.Database)
	if err != nil {
		log.LogFatal("Cannot connect to Database", err)
	}
	log.LogInfo("Database connection successful")

	err = grpc.RunGRPCService(store, &config, log)
	if err != nil {
		log.LogFatal("Cannot start server :", err)
		return err
	}
	return nil
}
