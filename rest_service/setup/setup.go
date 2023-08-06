package setup

import (
	"NOW/rest_service/db"
	"NOW/rest_service/logic/Handlers/report"
	"NOW/rest_service/logic/Handlers/user"
	Report "NOW/rest_service/logic/Repositories/report"
	User "NOW/rest_service/logic/Repositories/user"
	"NOW/rest_service/router"
	userRoutes "NOW/rest_service/router/routes"
)

func Routes(router router.Router, database *db.Database) {
	// Inicialización de los repositorios y los subrouters
	userRepo := User.NewUserRepositoryImpl(database)
	reportRepo := Report.NewReportRepositoryImplementation(database)

	userHandler := user.NewUserHandler(userRepo)
	reportHandler := report.NewReportHandler(reportRepo)

	userSubRouter := userRoutes.GetUserRoutes("/user", *userHandler)
	reportSubRouter := userRoutes.GetReportRoutes("/report", *reportHandler)

	// Añade el SubRouter al router
	router.AddSubRouter(userSubRouter)
	router.AddSubRouter(reportSubRouter)

	// Aquí puedes añadir rutas individuales que no requieren un subrouter
	// Ejemplo:
	// router.AddRoute(routingEntities.Route{
	// 	Method:     "GET",
	// 	Pattern:    "/health",
	// 	Handler:    HealthCheckHandler,
	// })
}
