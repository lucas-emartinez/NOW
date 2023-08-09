package main

import (
	"NOW/db"
	"NOW/router"
	"NOW/server"
	"NOW/setup"
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
