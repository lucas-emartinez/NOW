package router

import (
	RoutingEntity "NOW/rest_service/logic/entities/routing"
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	AddRoute(route RoutingEntity.Route)
	RouteHandler() http.HandlerFunc
	AddSubRouter(subRouter RoutingEntity.SubRouter)
}
