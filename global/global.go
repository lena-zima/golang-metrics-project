package global

import (
	"math/rand/v2"
	"runtime"
	"time"

	"github.com/lena-zima/golang-metrics-project/internal/storage/metricstorage"
)

var St metricstorage.MemStorage

var PollCount int64
var RandomValue = rand.Float64()
var RuntimeMetrics runtime.MemStats

const (
	PollInterval   time.Duration = 2
	ReportInterval time.Duration = 10
	ServerAddr     string        = "http://localhost:8080"
	ContentType    string        = "text/plain"
)

var RtMetrics = map[string]string{
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
