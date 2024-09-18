package main

import (
	"net/http"

	"github.com/lena-zima/golang-metrics-project/global"
	"github.com/lena-zima/golang-metrics-project/internal/server"
	"github.com/lena-zima/golang-metrics-project/internal/storage/metricstorage"
)

func main() {

	global.St.GaugeMetrics = make(map[string]metricstorage.Gauge, 0)
	global.St.CounterMetrics = make(map[string]metricstorage.Counter, 0)

	err := http.ListenAndServe(`:8080`, server.StartServer())
	if err != nil {
		panic(err)
	}
}
