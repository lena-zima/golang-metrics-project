package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type server struct {
	repo    repository.Repository
	addr    string
	handler *chi.Mux
}

func NewServer(config *serverconfig.ServerConfig, repo repository.Repository, handler *chi.Mux) (*server, error) {
	var serv server

	serv.repo = repo
	serv.addr = config.ServerAddr
	serv.handler = handler

	return &serv, nil
}

func (serv *server) RunJob() error {

	err := http.ListenAndServe(serv.addr, serv.handler)

	if err != nil {
		log.Printf("failed to get server config: %e", err)
		return err
	}

	return nil
}
