package main

import (
	"net/http"
	"strconv"
	"strings"
)

var storage MemStorage

// Supported metric types
type gauge float64
type counter int64

// Type to store metrics
type MemStorage struct {
	gaugeMetrics   map[string]gauge
	counterMetrics map[string]counter
}

func (storage *MemStorage) addGauge(name string, value float64) {
	storage.gaugeMetrics[name] = gauge(value)
}

func (storage *MemStorage) addCounter(name string, value int64) {
	storage.counterMetrics[name] += counter(value)
}



func updateHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		s := strings.Split(req.URL.Path, "/update/")
		metricInput := strings.Split(s[1], "/")

		if len(metricInput) != 3 {
			res.WriteHeader(http.StatusNotFound)
			return
		} else {
			metricType := metricInput[0]
			metricName := metricInput[1]
			metricValue := metricInput[2]

			switch metricType {
			case "gauge":
				value, err := strconv.ParseFloat(metricValue, 64)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					return
				}

				storage.addGauge(metricName, float64(value))

				res.WriteHeader(http.StatusOK)

			case "counter":
				value, err := strconv.Atoi(metricValue)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					return
				}

				storage.addCounter(metricName, int64(value))

				res.WriteHeader(http.StatusOK)

			default:
				res.WriteHeader(http.StatusBadRequest)
			}

		}
	}

}

func main() {

	storage.gaugeMetrics = make(map[string]gauge, 0)
	storage.counterMetrics = make(map[string]counter, 0)

	mux := http.NewServeMux()

	mux.HandleFunc("/update/", updateHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
