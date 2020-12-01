package core

import (
	"net/http"
)

// URLHandler Handler for representing url and corresponding handler
type URLHandler struct {
	URL         string
	Handler     http.HandlerFunc
	Methods     []string
	Middlewares []Middleware
}
