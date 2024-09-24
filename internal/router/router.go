package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/internal/handlers"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

func StartServer(repo repository.Repository) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostHandler(repo))

	return r
}
