package service

import (
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/grpc_api/user_service/db"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
)

type userService struct {
	*proto.UnimplementedUserServieServer
	log     logger.Logger
	storage *db.UserStore
}

// NewUserService creates a new gRPC server.
func NewUserService(storage *database.Storage, log logger.Logger) *userService {
	server := &userService{
		log: log,
	}
	userStore := db.NewUserStore(storage)
	server.storage = userStore
	return server
}
