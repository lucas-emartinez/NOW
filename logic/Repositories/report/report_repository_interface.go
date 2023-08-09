package report

import (
	dbEntity "NOW/logic/entities/db"
	"context"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *dbEntity.Report) error
}
