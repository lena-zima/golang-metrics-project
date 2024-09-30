package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

const (
	gauge   = "gauge"
	counter = "counter"
)

func GetAllHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gauges, counters, err := repo.GetAll()

		if err != nil {
			http.Error(w, "Failed to get metrics from store", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body><h1>Metrics</h1><ul>"))

		for k, v := range gauges {
			w.Write([]byte(fmt.Sprintf("<li>Gauge: %s = %g</li>", k, v)))
		}

		for k, v := range counters {
			w.Write([]byte(fmt.Sprintf("<li>Counter: %s = %d</li>", k, v)))
		}

		w.Write([]byte("</ul></body></html>"))

	}
}

func GetHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			metricType = chi.URLParam(r, "metricType")
			metricName = chi.URLParam(r, "metricName")
		)

		if metricType == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch metricType {
		case gauge:

			value, err := repo.GetGauge(metricName)

			if err != nil {
				http.Error(w, "Failed to get metric from store", http.StatusInternalServerError)
				return
			}

			if value == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%v", *value)))
			return

		case counter:

			value, err := repo.GetCounter(metricName)

			if err != nil {
				http.Error(w, "Failed to get metric from store", http.StatusInternalServerError)
				return
			}

			if value == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%v", *value)))
			return

		default:
			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

func PostHandler(repo repository.Repository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			metricType  = chi.URLParam(r, "metricType")
			metricName  = chi.URLParam(r, "metricName")
			metricValue = chi.URLParam(r, "metricValue")
		)

		if metricType != gauge && metricType != counter {
			http.Error(w, "Unknown metric type", http.StatusNotImplemented)
			return
		}

		if metricName == "" {
			http.Error(w, "Unknown metric name", http.StatusNotFound)
			return
		}

		if metricValue == "" {
			http.Error(w, "Empty metric value", http.StatusNotFound)
			return
		}

		switch metricType {
		case gauge:
			value, err := strconv.ParseFloat(metricValue, 64)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = repo.PostGauge(metricName, repository.Gauge(value))

			if err != nil {
				log.Printf("error while posting a gauge: %e", err)
			}

			w.WriteHeader(http.StatusOK)

		case counter:
			value, err := strconv.ParseInt(metricValue, 10, 64)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = repo.PostCounter(metricName, repository.Counter(value))

			if err != nil {
				log.Printf("error while posting a counter: %e", err)
			}

			w.WriteHeader(http.StatusOK)

		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}

}
