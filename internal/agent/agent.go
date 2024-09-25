package agent

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type gauge repository.Gauge
type counter repository.Counter

const (
	update             = "/update/"
	ContentType string = "text/plain"
)

type agent struct {
	metrics        *agentconfig.Metrics
	pollInterval   int
	reportInterval int
	serverAddr     string
}

func NewAgent(config *agentconfig.AgentConfig) (*agent, error) {
	var a agent

	a.metrics = config.Metrics
	a.pollInterval = config.PollInterval
	a.reportInterval = config.ReportInterval
	a.serverAddr = config.ServerAddr

	return &a, nil
}

func (a *agent) RunJob() {

	m := a.metrics

	//Variable which defines when to send metrics
	var sendCount int

	var reportCount = 5

	// Cycle to collect and send metrics
	for {
		time.Sleep(time.Second * time.Duration(a.pollInterval))

		err := collectMetrics(m)

		if err != nil {
			log.Printf("failed to get agent metrics: %e", err)
		}

		sendCount++

		if sendCount%reportCount == 0 {
			sendMetrics(m)
		}
	}
}

func collectMetrics(m *agentconfig.Metrics) error {

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	m.Alloc = repository.Gauge(rtm.Alloc)
	m.BuckHashSys = repository.Gauge(rtm.BuckHashSys)
	m.Frees = repository.Gauge(rtm.Frees)
	m.GCCPUFraction = repository.Gauge(rtm.GCCPUFraction)
	m.GCSys = repository.Gauge(rtm.GCSys)
	m.HeapAlloc = repository.Gauge(rtm.HeapAlloc)
	m.HeapIdle = repository.Gauge(rtm.HeapIdle)
	m.HeapInuse = repository.Gauge(rtm.HeapInuse)
	m.HeapObjects = repository.Gauge(rtm.HeapObjects)
	m.HeapReleased = repository.Gauge(rtm.HeapReleased)
	m.HeapSys = repository.Gauge(rtm.HeapSys)
	m.LastGC = repository.Gauge(rtm.LastGC)
	m.Lookups = repository.Gauge(rtm.Lookups)
	m.MCacheInuse = repository.Gauge(rtm.MCacheInuse)
	m.MCacheSys = repository.Gauge(rtm.MCacheSys)
	m.MSpanInuse = repository.Gauge(rtm.MSpanInuse)
	m.MSpanSys = repository.Gauge(rtm.MSpanSys)
	m.Mallocs = repository.Gauge(rtm.Mallocs)
	m.NextGC = repository.Gauge(rtm.NextGC)
	m.NumForcedGC = repository.Gauge(rtm.NumForcedGC)
	m.NumGC = repository.Gauge(rtm.NumGC)
	m.OtherSys = repository.Gauge(rtm.OtherSys)
	m.PauseTotalNs = repository.Gauge(rtm.PauseTotalNs)
	m.StackInuse = repository.Gauge(rtm.StackInuse)
	m.StackSys = repository.Gauge(rtm.StackSys)
	m.Sys = repository.Gauge(rtm.Sys)
	m.TotalAlloc = repository.Gauge(rtm.TotalAlloc)
	m.PollCount = m.PollCount + 1
	m.RandomValue = repository.Gauge(rand.Float64())

	return nil
}

func sendMetrics(m *agentconfig.Metrics) error {

	metrics := reflect.ValueOf(m)

	metricsType := reflect.TypeOf(m)

	for i := 0; i < metrics.Elem().NumField(); i++ {
		Mkey := reflect.ValueOf(metricsType.Elem().Field(i).Name)
		Mvalue := metrics.Elem().Field(i)
		Mtype := reflect.ValueOf(metrics.Elem().Field(i).Type().Name())

		err := sendMetric(Mtype, Mkey, Mvalue)

		if err != nil {
			log.Printf("failed to send metric: %e", err)
			return err
		}

	}

	return nil
}

func sendMetric(mtype reflect.Value, mname reflect.Value, mvalue reflect.Value) error {

	client := &http.Client{}

	url := fmt.Sprint("http://localhost:8080", update, mtype, "/", mname, "/", mvalue)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Printf("failed to send request to server: %e", err)
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("failed to send request to server: %e", err)
		return err
	}
	defer response.Body.Close()

	return nil
}
