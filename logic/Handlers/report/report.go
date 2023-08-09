package report

import (
	"NOW/logic/Repositories/report"
	dbEntity "NOW/logic/entities/db"
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

	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(err)

	err = h.repo.CreateReport(ctx, report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "No pudimos retransmitir la denuncia", http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(report)
}
