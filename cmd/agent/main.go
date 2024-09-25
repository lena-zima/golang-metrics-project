package main

import (
	"flag"
	"log"

	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {

	srvAddr := flag.String("a", "localhost:8080", "server endpoint address")
	repInt := flag.Int("r", 10, "report interval")
	pollInt := flag.Int("p", 2, "poll interval")

	flag.Parse()

	conf, err := agentconfig.GetConfig(*srvAddr, *repInt, *pollInt)

	if err != nil {
		log.Fatalf("failed to get agent config %e", err)
	}

	a, err := agent.NewAgent(conf)

	if err != nil {
		log.Fatalf("failed to create agent %e", err)
	}

	a.RunJob()

}
