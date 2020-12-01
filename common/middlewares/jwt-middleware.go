package middlewares

import (
	"context"
	"net/http"

	common ".."
	"../../core"
	"../../db/models"
	"github.com/robbert229/jwt"
)

var db = common.Dependencies.DB

// JWTMiddleware Middleware that attaches user object to request context
func JWTMiddleware() core.Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authorizationValues := r.Header.Values("Authorization")
			req := r
			if len(authorizationValues) > 0 {
				token := authorizationValues[0]
				algorithm := jwt.HmacSha256(common.JwtSecret)
				claims, err := algorithm.DecodeAndValidate(token)
				if err == nil {
					userID, err := claims.Get("userID")
					if err == nil {
						id := uint(userID.(float64))
						user := models.User{}
						db.First(&user, id)
						ctx := context.WithValue(req.Context(), "user", &user)
						req = r.WithContext(ctx)
					}
				}
			}
			f(w, req)
		}
	}
}
