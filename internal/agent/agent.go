package agent

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"strings"
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
	metrics        *Metrics
	pollInterval   int
	reportInterval int
	serverAddr     string
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

func NewAgent(config *agentconfig.AgentConfig) (*agent, error) {
	var a agent
	var err error

	a.metrics, err = initializeMetrics()

	if err != nil {
		log.Printf("error while metrics initialization: %e", err)
	}

	a.pollInterval = config.PollInterval
	a.reportInterval = config.ReportInterval
	a.serverAddr = config.ServerAddr

	return &a, err
}

func (a *agent) RunJob() {

	//Variable which defines when to send metrics
	var sendCount int

	if a.pollInterval == 0 {
		err := errors.New("polling interval cannot be 0")
		log.Printf("%e", err)
	}

	var reportCount = a.reportInterval / a.pollInterval

	// Cycle to collect and send metrics
	for {
		time.Sleep(time.Second * time.Duration(a.pollInterval))

		err := a.collectMetrics()

		if err != nil {
			log.Printf("failed to get agent metrics: %e", err)
		}

		sendCount++

		if reportCount == 0 {
			err := errors.New("reporting interval cannot be 0")
			log.Printf("%e", err)
		}

		if sendCount%reportCount == 0 {
			a.sendMetrics()
		}
	}
}

func initializeMetrics() (*Metrics, error) {
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

	return &m, nil
}

func (a *agent) collectMetrics() error {

	var m = a.metrics

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

func (a *agent) sendMetrics() error {

	var m = a.metrics

	metrics := reflect.ValueOf(m)

	metricsType := reflect.TypeOf(m)

	for i := 0; i < metrics.Elem().NumField(); i++ {
		Mkey := reflect.ValueOf(metricsType.Elem().Field(i).Name)
		Mvalue := metrics.Elem().Field(i)
		Mtype := reflect.ValueOf(metrics.Elem().Field(i).Type().Name()).String()
		Mtype = strings.ToLower(Mtype)

		err := a.sendMetric(Mtype, Mkey, Mvalue)

		if err != nil {
			log.Printf("failed to send metric: %e", err)
			return err
		}

	}

	return nil
}

func (a *agent) sendMetric(mtype string, mname reflect.Value, mvalue reflect.Value) error {

	client := &http.Client{}

	url := fmt.Sprintf("%s%s%s/%s/%v", a.serverAddr, update, mtype, mname, mvalue)

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
