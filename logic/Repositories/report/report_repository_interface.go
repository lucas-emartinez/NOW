package report

import (
	dbEntity "NOW/rest_service/logic/entities/db"
	"context"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *dbEntity.Report) error
}
