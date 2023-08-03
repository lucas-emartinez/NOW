package Routing

import (
	"NOW/db"
	"NOW/logic/Handlers"
	"NOW/logic/Repositories/user"
)

func RegisterRoutes(router RouterInterface, database *db.Database) {
	//Repositories usage
	userRepo := Repositories.NewUserRepositoryImpl(database)
	userHandler := Handlers.NewUserHandler(userRepo)

	router.AddRoute("/register", userHandler.CreateUser)
	router.AddRoute("/user/update", userHandler.UpdateUser)
	router.AddRoute("/user/{dni}", userHandler.GetByDNI)
}
