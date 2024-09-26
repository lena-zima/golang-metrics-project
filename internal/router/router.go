package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/handlers"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type server struct {
	repo    repository.Repository
	addr    string
	handler *chi.Mux
}

func NewServer(conf *serverconfig.ServerConfig) (*server, error) {

	var serv server
	serv.repo = conf.Repo
	serv.addr = conf.ServAddr

	r := chi.NewRouter()

	r.Get("/", handlers.GetAllHandler(serv.repo))
	r.Get("/value/{metricType}/{metricName}", handlers.GetHandler(serv.repo))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostHandler(serv.repo))

	serv.handler = r

	return &serv, nil
}

func (serv *server) StartServer() error {

	err := http.ListenAndServe(serv.addr, serv.handler)

	if err != nil {
		log.Printf("failed to get server config: %e", err)
		return err
	}

	return nil
}
