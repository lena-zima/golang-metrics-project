package main

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func main() {

	conf, err := serverconfig.GetConfig()

	if err != nil {
		log.Fatalf("failed to get server config %e", err)
	}

	r, err := router.NewServer(conf)

	if err != nil {
		log.Fatalf("failed to create a server %e", err)
	}

	router.StartServer(r)

}
