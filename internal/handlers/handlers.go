package handlers

import (
	"fmt"
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
		gauges, counters := repo.GetAll()

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
		} else {
			switch metricType {
			case gauge:

				value, exists := repo.GetGauge(metricName)

				if exists {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fmt.Sprintf("%v", value)))
				} else {
					w.WriteHeader(http.StatusNotFound)
				}

			case counter:

				value, exists := repo.GetCounter(metricName)

				if exists {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fmt.Sprintf("%v", value)))
				} else {
					w.WriteHeader(http.StatusNotFound)
				}

			default:
				w.WriteHeader(http.StatusBadRequest)
			}
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

		if metricType == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			switch metricType {
			case gauge:
				value, err := strconv.ParseFloat(metricValue, 64)

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				repo.PostGauge(metricName, repository.Gauge(value))

				w.WriteHeader(http.StatusOK)

			case counter:
				value, err := strconv.ParseInt(metricValue, 10, 64)

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				repo.PostCounter(metricName, repository.Counter(value))

				w.WriteHeader(http.StatusOK)

			default:
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

}
