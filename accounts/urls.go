package accounts

import (
	"../common/middlewares"
	"../core"
)

// URLPatterns Array of url handlers containing patterns and their corresponding handlers
var URLPatterns = []core.URLHandler{
	core.URLHandler{
		URL:     "/me",
		Handler: getUser,
		Methods: []string{"GET"},
		Middlewares: []core.Middleware{
			middlewares.AuthenticationMiddleware(),
		},
	},
	core.URLHandler{
		URL:     "/register",
		Handler: registerUser,
		Methods: []string{"POST"},
	},
	core.URLHandler{
		URL:     "/login",
		Handler: login,
		Methods: []string{"POST"},
	},
}
