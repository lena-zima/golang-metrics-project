package main

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
	"github.com/lena-zima/golang-metrics-project/internal/router"
	"github.com/lena-zima/golang-metrics-project/internal/server"
)

func main() {

	repo, err := memstorage.NewMemStorage()

	if err != nil {
		log.Printf("err while repo creation: %e", err)
	}

	conf, err := serverconfig.GetConfig()

	if err != nil {
		log.Fatalf("failed to get server config %e", err)
	}

	router, err := router.NewRouter(repo)

	if err != nil {
		log.Fatalf("failed to create a router %e", err)
	}

	serv, err := server.NewServer(conf, repo, router)

	if err != nil {
		log.Fatalf("failed to create a router %e", err)
	}

	serv.RunJob()

}
