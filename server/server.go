package server

import (
	"NOW/db"
	"NOW/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	listenAddr string
	Router     router.Router
	Database   *db.Database // Aquí asumimos que la base de datos también es una dependencia que se debe inyectar
}

func NewServer(database *db.Database, router router.Router) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Set a default port if there is nothing in the environment
	}
	return &Server{
		listenAddr: fmt.Sprintf(":%s", port),
		Router:     router,
		Database:   database,
	}
}

func (s *Server) Start() {
	// Antes de iniciar el servidor, configuras las rutas
	log.Printf("Server starting on port: %s", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.Router))
}
