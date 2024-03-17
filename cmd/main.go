package main

import (
	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/server"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
	_ "github.com/lib/pq"
)

func main() {
	log := logger.NewLogrusLogger("gestapo")
	config, err := config.LoadConfig("configs")
	if err != nil {
		log.LogFatal("Cannot load configuration:", err)
	}

	store, err := database.NewStorage(config.Database)
	if err != nil {
		log.LogFatal("Cannot connect to Database", err)
	}

	s3 := s3.NewS3Service(
		config.AwsS3.BucketName,
		config.AwsS3.Region,
		config.AwsS3.AccessKey,
		config.AwsS3.SecretKey,
	)

	server := server.NewServer(store, &config, log, s3)

	err = server.Start()
	if err != nil {
		log.LogFatal("Cannot start server :", err)
	}
}
