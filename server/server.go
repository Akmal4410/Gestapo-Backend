package server

import (
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/routes"
	"github.com/akmal4410/gestapo/util"
	"github.com/gorilla/mux"
)

// Server serves HTTP requests for our e-commerce service.
type Server struct {
	storage *database.Storage
	router  *mux.Router
	config  *util.Config
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(storage *database.Storage, config *util.Config) *Server {
	server := &Server{
		storage: storage,
		config:  config,
	}
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	router := mux.NewRouter()
	server.router = router

	server.setupRouter()

	fmt.Println("Go Bank Running on port :", server.config.ServerAddress)
	return http.ListenAndServe(server.config.ServerAddress, router)
}

func (server *Server) setupRouter() {
	routes.AuthRoutes(server.router, server.config)
}

type ErrorInfo struct {
	StatusCode int32  `json:"status_code"`
	Message    string `json:"message"`
}

// type ApiFunc func(http.ResponseWriter, *http.Request) error

// func MakeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := f(w, r); err != nil {
// 			util.WriteJSON(w, http.StatusBadRequest, ErrorInfo{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 		}
// 	}
// }

// func jsonContentTypeMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }
