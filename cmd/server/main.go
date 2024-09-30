package main

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/handlers"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
	"github.com/lena-zima/golang-metrics-project/internal/router"
	"github.com/lena-zima/golang-metrics-project/internal/server"
)

func main() {

	conf, err := serverconfig.GetConfig()

	if err != nil {
		log.Fatalf("failed to get server config %e", err)
	}

	repo, err := memstorage.NewMemStorage()

	if err != nil {
		log.Fatalf("err while repo creation: %e", err)
	}

	handler, err := handlers.NewHandler(repo)

	if err != nil {
		log.Fatalf("failed to create a handler %e", err)
	}

	router, err := router.NewRouter(handler)

	if err != nil {
		log.Fatalf("failed to create a router %e", err)
	}

	serverInstance, err := server.NewServer(conf, repo, router)

	if err != nil {
		log.Fatalf("failed to create a router %e", err)
	}

	serverInstance.RunJob()

}
