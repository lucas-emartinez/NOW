package Routing

import (
	"NOW/config"
	"NOW/db"
	"NOW/logic/Handlers"
	"NOW/logic/Repositories/user"
)

func RegisterRoutes(router RouterInterface, database *db.Database) {
	//Repositories usage
	userRepo := Repositories.NewUserRepositoryImpl(database)
	userHandler := Handlers.NewUserHandler(userRepo)

	router.AddRoute("/register", userHandler.CreateUser, config.VerifyNameMiddleware)
}
