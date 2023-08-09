package report

import (
	"NOW/db"
	dbEntity "NOW/logic/entities/db"
	"context"
	"strings"
)

type ReportRepositoryImplementation struct {
	db *db.Database
}

func NewReportRepositoryImplementation(db *db.Database) *ReportRepositoryImplementation {
	return &ReportRepositoryImplementation{db: db}
}

func (r *ReportRepositoryImplementation) CreateReport(ctx context.Context, report *dbEntity.Report) error {
	report.Coordinates = (strings.Replace(report.Coordinates, ",", " ", -1))
	report.Coordinates = "POINT(" + report.Coordinates + ")"
	sqlStatement := `INSERT INTO Report 
    				  		(ID, UserID, Description, Severity, Coordinates, Date, Timestamp)
					  VALUES (?, ?, ?, ?, ST_GeomFromText(?), ?, ?)`
	_, err := r.db.GetDB().ExecContext(ctx, sqlStatement, report.ID, report.UserID, report.Description, report.Severity, report.Coordinates, report.Date, report.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReportRepositoryImplementation) GetReport(id int) (*dbEntity.Report, error) {
	panic("implement me")
}

func (r *ReportRepositoryImplementation) DeleteReport() error {
	panic("implement me")
}

func (r *ReportRepositoryImplementation) GetAllReports() ([]*dbEntity.Report, error) {
	panic("implement me")
}
