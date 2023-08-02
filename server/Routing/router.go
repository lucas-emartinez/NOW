package Routing

import (
	"NOW/config"
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) AddRoute(pattern string, handler http.HandlerFunc, middleware ...config.Middleware) {
	// Aquí aplicaríamos los middlewares al handler
	for _, m := range middleware {
		handler = m(handler)
	}
	// Finalmente añadimos la ruta y el handler al multiplexor
	r.Mux.Handle(pattern, handler)
}

func (r *Router) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}
