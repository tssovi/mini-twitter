package main

import (
	"github.com/gorilla/mux"

	"./accounts"
	"./common/middlewares"
	"./core"
	"./tweets"
)

// projectRoutes Register all app routes here
var projectRoutes = []core.AppRoutes{
	core.AppRoutes{
		Prefix:      "/tweets",
		URLHandlers: tweets.URLPatterns,
		Middlewares: []core.Middleware{
			middlewares.AuthenticationMiddleware(),
			middlewares.JWTMiddleware(),
			middlewares.ResponseMiddleware(),
		},
	},
	core.AppRoutes{
		Prefix:      "/users",
		URLHandlers: accounts.URLPatterns,
		Middlewares: []core.Middleware{
			middlewares.JWTMiddleware(),
			middlewares.ResponseMiddleware(),
		},
	},
}

// GetProjectRouter Returns router containing all project routes
func GetProjectRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	core.SetAppRoutes(router, projectRoutes)
	return router
}
