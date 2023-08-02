package main

import (
	"NOW/db"
	"NOW/server"
	"NOW/server/Routing"
)

func main() {

	// Inicializa la base de datos
	database, err := db.NewDatabase()
	if err != nil {
		panic("Error")
		recover()
	}

	// Instanciacion del servidor
	app := server.NewServer()
	Routing.RegisterRoutes(app.Router, database)
	app.Start()
}
