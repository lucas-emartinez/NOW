package router

import (
	"NOW/rest_service/config"
	RoutingEntity "NOW/rest_service/logic/entities/routing"
	"context"
	"net/http"
	"regexp"
	"strings"
)

type RouterImplementation struct {
	Routes []RoutingEntity.Route
}

func NewRouter() Router {
	return &RouterImplementation{Routes: []RoutingEntity.Route{}}
}

func (r *RouterImplementation) AddRoute(route RoutingEntity.Route) {
	r.Routes = append(r.Routes, route)
}
func (r *RouterImplementation) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	config.LoggingMiddleware(r.RouteHandler()).ServeHTTP(w, req)
}
func (r *RouterImplementation) RouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, route := range r.Routes {
			if route.Method == req.Method {
				compiledPattern := regexp.MustCompile(route.Pattern)
				if matches := compiledPattern.FindStringSubmatch(req.URL.Path); matches != nil {
					req = addParamsToContext(req, route, matches[1:])
					applyMiddlewares(route.Middleware, route.Handler).ServeHTTP(w, req)
					return
				}
			}
		}
		http.NotFound(w, req)
	}
}
func (r *RouterImplementation) AddSubRouter(subRouter RoutingEntity.SubRouter) {
	for _, subRoute := range subRouter.Routes {
		pattern, params := CreateRegexPattern(subRouter.Prefix + subRoute.Pattern)
		r.Routes = append(r.Routes, RoutingEntity.Route{
			Method:     subRoute.Method,
			Pattern:    pattern,
			Params:     params,
			Handler:    subRoute.Handler,
			Middleware: subRoute.Middleware,
		})
	}
}

func CreateRegexPattern(pattern string) (string, []string) {
	parts := strings.Split(pattern, "/")
	params := []string{}

	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			params = append(params, part[1:len(part)-1])
			parts[i] = "([^/]+)"
		}
	}

	return strings.Join(parts, "/") + "$", params
}
func addParamsToContext(req *http.Request, route RoutingEntity.Route, matches []string) *http.Request {
	params := make(map[string]string, len(matches))
	for i, match := range matches {
		params[route.Params[i]] = match
	}
	return req.WithContext(context.WithValue(req.Context(), "params", params))
}
func applyMiddlewares(middlewares []config.Middleware, handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
