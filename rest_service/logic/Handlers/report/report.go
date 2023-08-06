package report

import (
	"NOW/rest_service/logic/Repositories/report"
	dbEntity "NOW/rest_service/logic/entities/db"
	"NOW/rest_service/messager"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ReportHandler struct {
	repo report.ReportRepository
}

func NewReportHandler(repo report.ReportRepository) *ReportHandler {
	return &ReportHandler{repo: repo}
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var report *dbEntity.Report
	log.Println(report)
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.CreateReport(ctx, report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(report)
	message := messager.WebsocketMessage{
		Type: "ReportCreated",
		Payload: map[string]interface{}{
			"reportID":    report.ID,
			"userID":      report.UserID,
			"description": report.Description,
			"severity":    report.Severity,
			"coordinates": report.Coordinates,
			"date":        report.Date,
			"timestamp":   report.Timestamp,
		},
	}

	err = messager.SendMessageToWebSocket(message)
	if err != nil {
		http.Error(w, "No pudimos retransmitir la denuncia", http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(report)
}
