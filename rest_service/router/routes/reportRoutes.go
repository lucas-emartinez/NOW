package routes

import (
	"NOW/rest_service/config"
	"NOW/rest_service/logic/Handlers/report"
	RoutingEntity "NOW/rest_service/logic/entities/routing"
	"net/http"
)

func GetReportRoutes(prefix string, reportHandler report.ReportHandler) RoutingEntity.SubRouter {
	subRouter := RoutingEntity.SubRouter{
		Prefix: prefix,
		Routes: []RoutingEntity.Route{
			{
				Method:     http.MethodPost,
				Pattern:    "/create",
				Handler:    reportHandler.CreateReport,
				Middleware: []config.Middleware{ /* middleware si es necesario */ },
			},
		},
	}

	return subRouter
}
