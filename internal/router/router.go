package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/internal/handlers"
)

func NewRouter(h *handlers.Handler) (*chi.Mux, error) {

	r := chi.NewRouter()

	r.Get("/", h.GetAllHandler)
	r.Get("/value/{metricType}/{metricName}", h.GetHandler)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.PostHandler)

	return r, nil
}
