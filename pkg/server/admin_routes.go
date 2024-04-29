package server

// func (server *Server) adminRoutes() {
// 	storage := db.NewAdminStore(server.storage)
// 	adminHandler := admin.NewAdminHandler(storage, server.log)

// 	tokenMaker, err := token.NewJWTMaker(server.config.TokenSymmetricKey)
// 	if err != nil {
// 		server.log.LogFatal("%w", err)
// 	}

// 	adminRoutes := server.router.PathPrefix("/admin").Subrouter()

// 	createCategory := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.CreateCategory))
// 	adminRoutes.Handle("/category", createCategory).Methods("POST")
// 	getCategory := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.GetCategories))
// 	adminRoutes.Handle("/category", getCategory).Methods("GET")

// 	allUsers := middleware.ApplyAccessRoleMiddleware(tokenMaker, server.log, utils.ADMIN, http.HandlerFunc(adminHandler.GetUsers))
// 	adminRoutes.Handle("/user", allUsers).Methods("GET")
// }
