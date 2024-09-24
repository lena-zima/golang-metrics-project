package main

import (
	"fmt"
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

	// Initiate metrics
	var m metric.Metrics

	// Initiate Config

	// Do agent

	var sendCount int

	fmt.Println("test")

	for {

		time.Sleep(pollInterval * time.Second)

		//fmt.Println("test")

		m.CollectMetrics()

		sendCount++

		if sendCount%reportCount == 0 {
			m.SendMetrics()
		}
	}

}
