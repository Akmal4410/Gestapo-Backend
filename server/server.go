package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/database"
	"github.com/gorilla/mux"
)

// Server serves HTTP requests for our e-commerce service.
type Server struct {
	listenAddress string
	storage       *database.Storage
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(listenAddress string, storage *database.Storage) *Server {
	server := &Server{
		listenAddress: listenAddress,
		storage:       storage,
	}
	return server
}

func (server *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/akmal", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("akmal")
	})

	fmt.Println("Go Bank Running on port", server.listenAddress)
	return http.ListenAndServe(server.listenAddress, router)

}
