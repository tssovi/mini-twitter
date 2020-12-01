package core

import "net/http"

// Middleware Type representing structure of a middleware function
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain Utility method for executing middlewares passed to it
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
