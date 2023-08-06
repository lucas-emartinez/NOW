package main

import (
	"NOW/rest_service/db"
	"NOW/rest_service/router"
	"NOW/rest_service/server"
	"NOW/rest_service/setup"
)

func main() {
	// Inicializa la base de datos
	database, err := db.NewDatabase()
	if err != nil {
		panic("Error")

	}

	// Router init
	router := router.NewRouter()

	// Add routes and subRouters to the Router
	setup.Routes(router, database)

	// Server init
	app := server.NewServer(database, router)

	app.Start()
}
