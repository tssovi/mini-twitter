package tweets

import (
	"../common/middlewares"
	"../core"
)

// URLPatterns Array of url handlers containing patterns and their corresponding handlers
var URLPatterns = []core.URLHandler{
	core.URLHandler{
		URL:     "/following",
		Handler: getFollowingUsers,
		Methods: []string{"GET", "POST"},
	},
	core.URLHandler{
		URL:     "/followers",
		Handler: getFollowers,
	},
	core.URLHandler{
		URL:     "/follow",
		Handler: followUser,
		Middlewares: []core.Middleware{
			middlewares.AuthenticationMiddleware(),
		},
		Methods: []string{"POST"},
	},
	core.URLHandler{
		URL:     "/tweet",
		Handler: createPost,
		Middlewares: []core.Middleware{
			middlewares.AuthenticationMiddleware(),
		},
		Methods: []string{"POST"},
	},
	core.URLHandler{
		URL:     "/feed",
		Handler: getFeed,
		Middlewares: []core.Middleware{
			middlewares.AuthenticationMiddleware(),
		},
		Methods: []string{"GET"},
	},
}
