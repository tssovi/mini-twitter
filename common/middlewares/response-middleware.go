package middlewares

import (
	"net/http"

	"../../core"
)

// ResponseMiddleware Middleware that attaches Content-Type header to response writer
func ResponseMiddleware() core.Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			f(w, r)
		}
	}
}
