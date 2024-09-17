package main

import (
	"net/http"

	"github.com/lena-zima/golang-metrics-project/global"
	"github.com/lena-zima/golang-metrics-project/internal/server"
	"github.com/lena-zima/golang-metrics-project/internal/storage/metrics_storage"
)

func main() {

	global.St.GaugeMetrics = make(map[string]metrics_storage.Gauge, 0)
	global.St.CounterMetrics = make(map[string]metrics_storage.Counter, 0)

	err := http.ListenAndServe(`:8080`, server.StartServer())
	if err != nil {
		panic(err)
	}
}
