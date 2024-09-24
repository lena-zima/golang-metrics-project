package main

import (
	"time"

	"github.com/lena-zima/golang-metrics-project/internal/metric"
)

const (
	pollInterval   time.Duration = 2
	reportInterval time.Duration = 10
	reportCount    int           = 5
	ServerAddr     string        = "http://localhost:8080"
	ContentType    string        = "text/plain"
)

func main() {

	// Step 1. Initiate metrics
	var m metric.Metrics

	// Step 2. Initiate Config
	// TODO

	// Step 3. Operate metrics

	// Variable which defines when to send metrics
	var sendCount int

	// Cycle to collect and send metrics
	for {

		time.Sleep(pollInterval * time.Second)

		m.CollectMetrics()

		sendCount++

		if sendCount%reportCount == 0 {
			m.SendMetrics()
		}
	}

}
