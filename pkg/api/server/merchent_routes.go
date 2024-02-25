package server

import "github.com/akmal4410/gestapo/pkg/api/merchent"

type MerchentRoute struct {
	merchent *merchent.MerchentHandler
}

var mechentRoute MerchentRoute

func (server *Server) merchentRoutes() {
	mechentRoute.merchent = merchent.NewMerchentHandler()

	merchentRoutes := server.router.PathPrefix("/merchent").Subrouter()

	merchentRoutes.HandleFunc("/profile/{id}", mechentRoute.merchent.ViewProfile).Methods("GET")
}
