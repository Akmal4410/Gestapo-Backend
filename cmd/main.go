package main

import (
	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/server"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	_ "github.com/lib/pq"
)

func main() {
	log := logger.NewLogrusLogger("gestapo")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.LogFatal("Cannot load configuration:", err)
	}

	store, err := database.NewStorage(config.Database)
	if err != nil {
		log.LogFatal("Cannot connect to Database", err)
	}

	server := server.NewServer(store, &config, log)

	err = server.Start()
	if err != nil {
		log.LogFatal("Cannot start server :", err)
	}
}
