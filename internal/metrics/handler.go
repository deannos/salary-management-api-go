package metrics

import (
	"database/sql"
	"net/http"
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
	w.WriteHeader(200)
}
