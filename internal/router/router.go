package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/handlers"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type server struct {
	repo repository.Repository
}

func StartServer(conf *serverconfig.ServerConfig) *chi.Mux {

	var serv server
	serv.repo = conf.Repo

	r := chi.NewRouter()

	r.Get("/", handlers.GetAllHandler(serv.repo))
	r.Get("/value/{metricType}/{metricName}", handlers.GetHandler(serv.repo))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostHandler(serv.repo))

	return r
}
