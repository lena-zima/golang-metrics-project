package main

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {
	conf, err := agentconfig.GetConfig()

	if err != nil {
		log.Fatalf("failed to get agent config %e", err)
	}

	a, err := agent.NewAgent(conf)

	if err != nil {
		log.Fatalf("failed to create agent %e", err)
	}

	a.RunJob()

}
