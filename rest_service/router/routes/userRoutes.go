package routes

import (
	"NOW/rest_service/config"
	"NOW/rest_service/logic/Handlers/user"
	RoutingEntity "NOW/rest_service/logic/entities/routing"
	"net/http"
)

func GetUserRoutes(prefix string, userHandler user.UserHandler) RoutingEntity.SubRouter {
	subRouter := RoutingEntity.SubRouter{
		Prefix: prefix,
		Routes: []RoutingEntity.Route{
			{
				Method:     http.MethodPost,
				Pattern:    "/register",
				Handler:    userHandler.CreateUser,
				Middleware: []config.Middleware{ /* middleware si es necesario */ },
			},
			{
				Method:     http.MethodGet,
				Pattern:    "/{dni}",
				Handler:    userHandler.GetByDNI,
				Middleware: []config.Middleware{ /* middleware si es necesario */ },
			},
			{
				Method:     http.MethodDelete,
				Pattern:    "/{dni}",
				Handler:    userHandler.DeleteUser,
				Middleware: []config.Middleware{ /* middleware si es necesario */ },
			},
			{
				Method:     http.MethodPut,
				Pattern:    "/update/{dni}",
				Handler:    userHandler.UpdateUser,
				Middleware: []config.Middleware{ /* middleware si es necesario */ },
			},

			// Otras rutas del usuario...
		},
	}

	return subRouter
}
