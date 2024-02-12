package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/akmal4410/gestapo/util"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Working on Gestapo..")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}

	db, err := sql.Open(config.DBServer, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot verify a connection:", err)
	} else {
		fmt.Println("Successfully connected to the database.")
	}

}
