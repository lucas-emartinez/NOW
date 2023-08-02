package Routing

import (
	"NOW/config"
	"net/http"
)

type RouterInterface interface {
	// AddRoute adds a route to the router
	AddRoute(pattern string, handler http.HandlerFunc, middleware ...config.Middleware)
}
