package metric

import (
	"fmt"
	"math/rand"
	"net/http"

	//"net/http"
	"reflect"
	"runtime"
)

type gauge float64
type counter int64

type Metrics struct {
	//Runtime
	Alloc         gauge
	BuckHashSys   gauge
	Frees         gauge
	GCCPUFraction gauge
	GCSys         gauge
	HeapAlloc     gauge
	HeapIdle      gauge
	HeapInuse     gauge
	HeapObjects   gauge
	HeapReleased  gauge
	HeapSys       gauge
	LastGC        gauge
	Lookups       gauge
	MCacheInuse   gauge
	MCacheSys     gauge
	MSpanInuse    gauge
	MSpanSys      gauge
	Mallocs       gauge
	NextGC        gauge
	NumForcedGC   gauge
	NumGC         gauge
	OtherSys      gauge
	PauseTotalNs  gauge
	StackInuse    gauge
	StackSys      gauge
	Sys           gauge
	TotalAlloc    gauge

	// Custom
	PollCount   counter
	RandomValue gauge
}

func (m *Metrics) CollectMetrics() error {

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	m.Alloc = gauge(rtm.Alloc)
	m.BuckHashSys = gauge(rtm.BuckHashSys)
	m.Frees = gauge(rtm.Frees)
	m.GCCPUFraction = gauge(rtm.GCCPUFraction)
	m.GCSys = gauge(rtm.GCSys)
	m.HeapAlloc = gauge(rtm.HeapAlloc)
	m.HeapIdle = gauge(rtm.HeapIdle)
	m.HeapInuse = gauge(rtm.HeapInuse)
	m.HeapObjects = gauge(rtm.HeapObjects)
	m.HeapReleased = gauge(rtm.HeapReleased)
	m.HeapSys = gauge(rtm.HeapSys)
	m.LastGC = gauge(rtm.LastGC)
	m.Lookups = gauge(rtm.Lookups)
	m.MCacheInuse = gauge(rtm.MCacheInuse)
	m.MCacheSys = gauge(rtm.MCacheSys)
	m.MSpanInuse = gauge(rtm.MSpanInuse)
	m.MSpanSys = gauge(rtm.MSpanSys)
	m.Mallocs = gauge(rtm.Mallocs)
	m.NextGC = gauge(rtm.NextGC)
	m.NumForcedGC = gauge(rtm.NumForcedGC)
	m.NumGC = gauge(rtm.NumGC)
	m.OtherSys = gauge(rtm.OtherSys)
	m.PauseTotalNs = gauge(rtm.PauseTotalNs)
	m.StackInuse = gauge(rtm.StackInuse)
	m.StackSys = gauge(rtm.StackSys)
	m.Sys = gauge(rtm.Sys)
	m.TotalAlloc = gauge(rtm.TotalAlloc)
	m.PollCount++
	m.RandomValue = gauge(rand.Float64())

	return nil
}

func (m *Metrics) SendMetrics() {

	metrics := reflect.ValueOf(m)

	metricsType := reflect.TypeOf(m)

	for i := 0; i < metrics.Elem().NumField(); i++ {
		Mkey := reflect.ValueOf(metricsType.Elem().Field(i).Name)
		Mvalue := reflect.ValueOf(metrics.Elem().Field(i))
		if Mkey == reflect.ValueOf("PollCount") {
			sendMetric("counter", Mkey, Mvalue)
		} else {
			sendMetric("gauge", Mkey, Mvalue)
		}
	}
}

func sendMetric(mtype string, mname reflect.Value, mvalue reflect.Value) {

	client := &http.Client{}

	url := fmt.Sprint("http://localhost:8080/update/", mtype, "/", mname, "/", mvalue)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
}
