package router

import (
	"NOW/logic/entities/routing"
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	AddRoute(route routing.Route)
	RouteHandler() http.HandlerFunc
	AddSubRouter(subRouter routing.SubRouter)
}
