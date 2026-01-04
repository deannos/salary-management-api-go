package metrics

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		service: NewService(db),
	}
}

func (h *Handler) GetCountryMetrics(w http.ResponseWriter, r *http.Request) {
	country := strings.TrimPrefix(r.URL.Path, "/metrics/country/")

	result, err := h.service.ByCountry(country)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"min": result.Min,
		"max": result.Max,
		"avg": result.Avg,
	})
}

func (h *Handler) GetJobTitleMetrics(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/metrics/job-title/")

	avg, err := h.service.AverageByJobTitle(title)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"avg": avg,
	})
}
