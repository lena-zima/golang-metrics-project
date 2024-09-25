package main

import (
	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {
	conf := agentconfig.GetConfig()

	a, err := agent.NewAgent(conf)

	a.RunJob()

	if err != nil {
		panic("AA")
	}

}
