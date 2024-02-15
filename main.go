package main

import (
	"log"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/server"
	"github.com/akmal4410/gestapo/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}

	store, err := database.NewStorage(config)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}

	server := server.NewServer(store, &config)

	err = server.Start()
	if err != nil {
		log.Fatal("Cannot start server :", err)
	}
}
