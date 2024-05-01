package server

import (
	"net/http"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
	"github.com/gorilla/mux"
)

// Server serves HTTP requests for our e-commerce service.
type Server struct {
	storage *database.Storage
	router  *mux.Router
	config  *config.Config
	log     logger.Logger
	s3      *s3.S3Service
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(storage *database.Storage, config *config.Config, log logger.Logger, s3 *s3.S3Service) *Server {
	server := &Server{
		storage: storage,
		config:  config,
		log:     log,
		s3:      s3,
	}
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	router := mux.NewRouter()
	server.router = router

	// helpers.RegisterValidator()

	server.setupRouter()
	server.log.LogInfo("Go Bank Running on port :", server.config.ServerAddress)
	return http.ListenAndServe("", server.router)
}

func (server *Server) setupRouter() {
	// server.merchantRoutes()
	// server.userRoutes()
}
