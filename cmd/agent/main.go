package main

import (
	"fmt"
	"log"

	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {

	conf, err := agentconfig.GetConfig()

	fmt.Print(conf)

	if err != nil {
		log.Fatalf("failed to get agent config %e", err)
	}

	agentInstance, err := agent.NewAgent(conf)

	if err != nil {
		log.Fatalf("failed to create agent %e", err)
	}

	agentInstance.RunJob()

}
