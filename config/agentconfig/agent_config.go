package agentconfig

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

type gauge repository.Gauge
type counter repository.Counter

type AgentConfig struct {
	Metrics        *Metrics
	PollInterval   int
	ReportInterval int
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

func GetConfig() (*AgentConfig, error) {

	var conf AgentConfig

	conf.Metrics = initializeMetrics()

	//Server Address
	srvAddr, err := getEnv("srvAddr")

	if err != nil {
		log.Printf("failed to get server address env: %e", err)
	}

	conf.ServerAddr = "http://" + srvAddr

	if srvAddr == "" {
		srvAddr, err = getFlag("srvAddr")
		conf.ServerAddr = "http://" + srvAddr

		if err != nil {
			log.Printf("failed to get server address flag: %e", err)
		}
	}

	// pollInt

	pollInt, err := getEnv("pollInt")

	fmt.Printf(pollInt)

	if err != nil {
		log.Printf("failed to get poll interval env: %e", err)
	}

	if pollInt == "" {
		pollInt, err = getFlag("pollInt")

		if err != nil {
			log.Printf("failed to get poll interval flag: %e", err)
		}

		conf.PollInterval, err = strconv.Atoi(pollInt)

		if err != nil {
			log.Printf("failed to get poll interval flag: %e", err)
		}
	} else {
		conf.PollInterval, err = strconv.Atoi(pollInt)

		if err != nil {
			log.Printf("failed to get poll interval env: %e", err)
		}
	}

	//repInt

	repInt, err := getEnv("repInt")

	if err != nil {
		log.Printf("failed to get rep interval env: %e", err)
	}

	if repInt == "" {
		repInt, err = getFlag("repInt")

		if err != nil {
			log.Printf("failed to get poll interval flag: %e", err)
		}

		conf.ReportInterval, err = strconv.Atoi(repInt)

		if err != nil {
			log.Printf("failed to get poll interval flag: %e", err)
		}
	} else {
		conf.ReportInterval, err = strconv.Atoi(repInt)

		if err != nil {
			log.Printf("failed to get rep interval env: %e", err)
		}
	}

	return &conf, nil
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

func getEnv(name string) (string, error) {

	switch name {
	case "srvAddr":
		srvEnv, srvEnvExists := os.LookupEnv("ADDRESS")
		if srvEnvExists {
			return srvEnv, nil
		}
	case "pollInt":
		repEnv, repEnvExists := os.LookupEnv("REPORT_INTERVAL")

		if repEnvExists {
			return repEnv, nil
		}
	case "repInt":
		pollEnv, pollEnvExists := os.LookupEnv("POLL_INTERVAL")

		if pollEnvExists {
			return pollEnv, nil
		}
	default:
		err := errors.New("unknown env")
		log.Printf("env parsing error: %e", err)
		return "", err
	}
	return "", nil
}

func getFlag(name string) (string, error) {
	switch name {
	case "srvAddr":
		srvAddr := flag.String("a", "localhost:8080", "server endpoint address")
		flag.Parse()
		return *srvAddr, nil
	case "pollInt":
		pollInt := flag.String("p", "2", "poll interval")
		flag.Parse()
		return *pollInt, nil
	case "repInt":
		repInt := flag.String("r", "10", "report interval")
		flag.Parse()
		return *repInt, nil
	default:
		err := errors.New("unknown flag")
		log.Printf("env parsing flags: %e", err)
		return "", err
	}

}
