package routing

import (
	"NOW/config"
	"net/http"
)

type Route struct {
	Pattern    string
	Params     []string
	Handler    http.HandlerFunc
	Middleware []config.Middleware
	Method     string
}
