package service

import (
	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
)

// AuthenticationService serves gRPC requests for our e-commerce service.
type AuthenticationService struct {
	proto.UnimplementedAuthenticationServiceServer
	storage *database.Storage
	config  *config.Config
	log     logger.Logger
	s3      *s3.S3Service
}

// NewAuthenticationService creates a new gRPC server and sets up routing.
func NewAuthenticationService(storage *database.Storage, config *config.Config, log logger.Logger) *AuthenticationService {
	server := &AuthenticationService{
		storage: storage,
		config:  config,
		log:     log,
	}
	s3 := s3.NewS3Service(
		config.AwsS3.BucketName,
		config.AwsS3.Region,
		config.AwsS3.AccessKey,
		config.AwsS3.SecretKey,
	)
	server.s3 = s3
	return server
}
