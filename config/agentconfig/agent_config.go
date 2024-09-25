package agentconfig

import (
	"time"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type gauge repository.Gauge
type counter repository.Counter

type AgentConfig struct {
	Metrics        *Metrics
	PollInterval   time.Duration
	ReportInterval time.Duration
	ServerAddr     string
}

type Metrics struct {
	//Runtime
	Alloc         repository.Gauge
	BuckHashSys   repository.Gauge
	Frees         repository.Gauge
	GCCPUFraction repository.Gauge
	GCSys         repository.Gauge
	HeapAlloc     repository.Gauge
	HeapIdle      repository.Gauge
	HeapInuse     repository.Gauge
	HeapObjects   repository.Gauge
	HeapReleased  repository.Gauge
	HeapSys       repository.Gauge
	LastGC        repository.Gauge
	Lookups       repository.Gauge
	MCacheInuse   repository.Gauge
	MCacheSys     repository.Gauge
	MSpanInuse    repository.Gauge
	MSpanSys      repository.Gauge
	Mallocs       repository.Gauge
	NextGC        repository.Gauge
	NumForcedGC   repository.Gauge
	NumGC         repository.Gauge
	OtherSys      repository.Gauge
	PauseTotalNs  repository.Gauge
	StackInuse    repository.Gauge
	StackSys      repository.Gauge
	Sys           repository.Gauge
	TotalAlloc    repository.Gauge

	// Custom
	PollCount   repository.Counter
	RandomValue repository.Gauge
}

func GetConfig() *AgentConfig {

	var conf AgentConfig
	conf.Metrics = initializeMetrics()
	conf.PollInterval = 2
	conf.ReportInterval = 10
	conf.ServerAddr = "http://localhost:8080"

	return &conf
}

func initializeMetrics() *Metrics {
	var m Metrics

	m.Alloc = repository.Gauge(0)
	m.BuckHashSys = repository.Gauge(0)
	m.Frees = repository.Gauge(0)
	m.GCCPUFraction = repository.Gauge(0)
	m.GCSys = repository.Gauge(0)
	m.HeapAlloc = repository.Gauge(0)
	m.HeapIdle = repository.Gauge(0)
	m.HeapInuse = repository.Gauge(0)
	m.HeapObjects = repository.Gauge(0)
	m.HeapReleased = repository.Gauge(0)
	m.HeapSys = repository.Gauge(0)
	m.LastGC = repository.Gauge(0)
	m.Lookups = repository.Gauge(0)
	m.MCacheInuse = repository.Gauge(0)
	m.MCacheSys = repository.Gauge(0)
	m.MSpanInuse = repository.Gauge(0)
	m.MSpanSys = repository.Gauge(0)
	m.Mallocs = repository.Gauge(0)
	m.NextGC = repository.Gauge(0)
	m.NumForcedGC = repository.Gauge(0)
	m.NumGC = repository.Gauge(0)
	m.OtherSys = repository.Gauge(0)
	m.PauseTotalNs = repository.Gauge(0)
	m.StackInuse = repository.Gauge(0)
	m.StackSys = repository.Gauge(0)
	m.Sys = repository.Gauge(0)
	m.TotalAlloc = repository.Gauge(0)

	// Custom
	m.PollCount = repository.Counter(0)
	m.RandomValue = repository.Gauge(0)

	return &m
}
