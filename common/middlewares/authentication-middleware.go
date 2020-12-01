package middlewares

import (
	"net/http"

	"../../core"
)

// AuthenticationMiddleware Middleware that validates the request user, else returns forbidden
func AuthenticationMiddleware() core.Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user := ctx.Value("user")
			if user == nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			f(w, r)
		}
	}
}
