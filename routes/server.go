package routes

import (
	"net/http"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/services/logger"
	"github.com/akmal4410/gestapo/utils"
	"github.com/gorilla/mux"
)

// Server serves HTTP requests for our e-commerce service.
type Server struct {
	storage *database.Storage
	router  *mux.Router
	config  *utils.Config
	log     logger.Logger
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(storage *database.Storage, config *utils.Config, log logger.Logger) *Server {
	server := &Server{
		storage: storage,
		config:  config,
		log:     log,
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
	return http.ListenAndServe(server.config.ServerAddress, router)
}

func (server *Server) setupRouter() {
	server.authRoutes()
}
