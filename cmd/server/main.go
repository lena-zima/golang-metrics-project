package main

import (
	"flag"
	"log"
	"os"

	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func main() {

	srvAddr := flag.String("a", "localhost:8080", "server endpoint address")

	flag.Parse()

	srvEnv, srvEnvExists := os.LookupEnv("ADDRESS")

	if srvEnvExists == true {
		*srvAddr = srvEnv
	}

	conf, err := serverconfig.GetConfig(*srvAddr)

	if err != nil {
		log.Fatalf("failed to get server config %e", err)
	}

	serv, err := router.NewServer(conf)

	if err != nil {
		log.Fatalf("failed to create a server %e", err)
	}

	serv.StartServer()

}
