package core

import (
	"github.com/gorilla/mux"
)

// AppRoutes Type representing an app's url prefix and handlers
type AppRoutes struct {
	Prefix      string
	URLHandlers []URLHandler
	Middlewares []Middleware
}

// SetAppRoutes Utility method for attaching app routes to provided router
func SetAppRoutes(r *mux.Router, projectRoutes []AppRoutes) {
	for _, appRoute := range projectRoutes {
		s := r.PathPrefix(appRoute.Prefix).Subrouter()
		for _, handler := range appRoute.URLHandlers {
			middlewares := []Middleware{}
			if len(handler.Middlewares) > 0 {
				middlewares = append(middlewares, handler.Middlewares...)
			}
			if len(appRoute.Middlewares) > 0 {
				middlewares = append(middlewares, appRoute.Middlewares...)
			}
			route := s.HandleFunc(handler.URL, Chain(handler.Handler, middlewares...))
			if len(handler.Methods) > 0 {
				route.Methods(handler.Methods...)
			}
		}
	}
}
