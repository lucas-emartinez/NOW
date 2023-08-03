package Routing

import (
	"NOW/config"
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	return &Router{
		Mux: mux,
	}
}

func extractRouteParameters(pattern string, req *http.Request) ([]string, error) {
	segments := strings.Split(req.URL.Path, "/")
	return segments, nil
}

func (r *Router) AddRoute(pattern string, handler http.HandlerFunc, middleware ...config.Middleware) {

	// Aquí aplicaríamos los middlewares al handler
	for _, m := range middleware {
		handler = m(handler)
	}
	// Finalmente añadimos la ruta y el handler al multiplexor
	r.Mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		segments, err := extractRouteParameters(pattern, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(segments)
		// Ahora "segments" es un slice con todos los parámetros de la ruta.
		// Puedes usar este slice para acceder a los parámetros de ruta.
		handler(w, req)
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := config.LoggingMiddleware(r.Mux.ServeHTTP)
	handler(w, req)
}
