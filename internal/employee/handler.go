package employee

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
