package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type server struct {
	repo   repository.Repository
	addr   string
	router *chi.Mux
}

func NewServer(config *serverconfig.ServerConfig, repo repository.Repository, router *chi.Mux) (*server, error) {
	var serv server

	serv.repo = repo
	serv.addr = config.ServerAddr
	serv.router = router

	return &serv, nil
}

func (serv *server) RunJob() error {

	err := http.ListenAndServe(serv.addr, serv.router)

	if err != nil {
		log.Printf("failed to get server config: %e", err)
		return err
	}

	return nil
}
