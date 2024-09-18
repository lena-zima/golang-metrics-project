package main

import (
	"time"

	"github.com/lena-zima/golang-metrics-project/global"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {

	for {
		time.Sleep(global.PollInterval * time.Second)

		agent.UpdateMetrics()

		agent.SendMetrics()

	}

}
