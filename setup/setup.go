package setup

import (
	"NOW/db"
	"NOW/logic/Handlers/report"
	"NOW/logic/Handlers/user"
	Report "NOW/logic/Repositories/report"
	User "NOW/logic/Repositories/user"
	"NOW/router"
	"NOW/router/routes"
)

func Routes(router router.Router, database *db.Database) {
	// Inicialización de los repositorios y los subrouters
	userRepo := User.NewUserRepositoryImpl(database)
	reportRepo := Report.NewReportRepositoryImplementation(database)

	userHandler := user.NewUserHandler(userRepo)
	reportHandler := report.NewReportHandler(reportRepo)

	userSubRouter := routes.GetUserRoutes("/user", *userHandler)
	reportSubRouter := routes.GetReportRoutes("/report", *reportHandler)

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
