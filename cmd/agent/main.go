package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

var rtMetrics = map[string]string{
	"Alloc":         "gauge",
	"BuckHashSys":   "gauge",
	"Frees":         "gauge",
	"GCCPUFraction": "gauge",
	"GCSys":         "gauge",
	"HeapAlloc":     "gauge",
	"HeapIdle":      "gauge",
	"HeapInuse":     "gauge",
	"HeapObjects":   "gauge",
	"HeapReleased":  "gauge",
	"HeapSys":       "gauge",
	"LastGC":        "gauge",
	"Lookups":       "gauge",
	"MCacheInuse":   "gauge",
	"MCacheSys":     "gauge",
	"MSpanInuse":    "gauge",
	"MSpanSys":      "gauge",
	"Mallocs":       "gauge",
	"NextGC":        "gauge",
	"NumForcedGC":   "gauge",
	"NumGC":         "gauge",
	"OtherSys":      "gauge",
	"PauseTotalNs":  "gauge",
	"StackInuse":    "gauge",
	"StackSys":      "gauge",
	"Sys":           "gauge",
	"TotalAlloc":    "gauge",
}

var PollCount int64
var RandomValue = rand.Float64()
var RuntimeMetrics runtime.MemStats

const (
	pollInterval   time.Duration = 2
	reportInterval time.Duration = 10
	serverAddr     string        = "http://localhost:8080"
	contentType    string        = "text/plain"
)

func updateMetrics() {
	runtime.ReadMemStats(&RuntimeMetrics)
	PollCount++
	RandomValue = rand.Float64()
}

func SendMetrics() {
	for k, v := range rtMetrics {
		val := reflect.ValueOf(RuntimeMetrics).FieldByName(k)
		sendMetric(v, k, val)
	}

	PollCountvalue := reflect.ValueOf(PollCount)
	sendMetric("counter", "PollCount", PollCountvalue)

	RandomValuevalue := reflect.ValueOf(RandomValue)
	sendMetric("gauge", "RandomValue", RandomValuevalue)
}

func sendMetric(mtype string, mname string, mvalue reflect.Value) {

	client := &http.Client{}

	url := fmt.Sprint(serverAddr, "/update/", mtype, "/", mname, "/", mvalue)

	//fmt.Println(url)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		panic(err)
	}
	client.Do(request)

}

func main() {

	for true {
		time.Sleep(pollInterval * time.Second)

		updateMetrics()

		SendMetrics()

	}

}
