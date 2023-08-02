package server

import (
	"NOW/server/Routing"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	listenAddr string
	Router     *Routing.Router
}

func NewServer() *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Set a default port if there is nothing in the environment
	}
	return &Server{
		listenAddr: fmt.Sprintf(":%s", port),
		Router:     Routing.NewRouter(),
	}
}

func (s *Server) Start() {
	log.Printf("Server starting in port: %s", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.Router.Mux))
}
