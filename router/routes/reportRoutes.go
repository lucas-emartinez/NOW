package routes

import (
	"NOW/config"
	"NOW/logic/Handlers/report"
	"NOW/logic/entities/routing"
	"net/http"
)

func GetReportRoutes(prefix string, reportHandler report.ReportHandler) routing.SubRouter {
	subRouter := routing.SubRouter{
		Prefix: prefix,
		Routes: []routing.Route{
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
