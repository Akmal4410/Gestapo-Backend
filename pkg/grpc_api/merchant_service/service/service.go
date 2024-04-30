package service

import (
	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/authentication_service/db"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
)

// merchantService serves gRPC requests for our e-commerce service.
type merchantService struct {
	proto.UnimplementedMerchantServiceServer
	config  *config.Config
	log     logger.Logger
	s3      *s3.S3Service
	storage *db.AuthStore
	token   token.Maker
}

// NewMerchantService creates a new gRPC server.
func NewMerchantService(storage *database.Storage, config *config.Config, log logger.Logger, tokenMaker token.Maker) *merchantService {
	server := &merchantService{
		config: config,
		log:    log,
		token:  tokenMaker,
	}
	s3 := s3.NewS3Service(
		config.AwsS3.BucketName,
		config.AwsS3.Region,
		config.AwsS3.AccessKey,
		config.AwsS3.SecretKey,
	)

	authStore := db.NewAuthStore(storage)

	server.s3 = s3
	server.storage = authStore
	return server
}