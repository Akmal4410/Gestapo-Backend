package main

import (
	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/routes"
	"github.com/akmal4410/gestapo/services/logger"
	"github.com/akmal4410/gestapo/utils"
	_ "github.com/lib/pq"
)

func main() {
	log := logger.NewLogrusLogger("gestapo")
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.LogFatal("Cannot load configuration:", err)
	}

	store, err := database.NewStorage(config.Database)
	if err != nil {
		log.LogFatal("Cannot connect to Database", err)
	}

	server := routes.NewServer(store, &config, log)

	err = server.Start()
	if err != nil {
		log.LogFatal("Cannot start server :", err)
	}
}
