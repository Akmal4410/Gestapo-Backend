package main

import (
	"log"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/server"
	"github.com/akmal4410/gestapo/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}

	store, err := database.NewStorage(config.Database)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}

	server := server.NewServer(store, &config)

	err = server.Start()
	if err != nil {
		log.Fatal("Cannot start server :", err)
	}
}
