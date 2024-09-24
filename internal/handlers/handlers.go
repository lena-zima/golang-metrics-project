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

func GetHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse URL
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

				value := repo.GetGauge(metricName)

				w.WriteHeader(http.StatusOK)

				w.Write([]byte(fmt.Sprintf("%v", value)))

			case counter:

				value := repo.GetCounter(metricName)

				w.WriteHeader(http.StatusOK)

				w.Write([]byte(fmt.Sprintf("%v", value)))

			default:
				w.WriteHeader(http.StatusBadRequest)
			}
		}

	}
}

func PostHandler(repo repository.Repository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Parse URL
		var (
			metricType  = chi.URLParam(r, "metricType")
			metricName  = chi.URLParam(r, "metricName")
			metricValue = chi.URLParam(r, "metricValue")
		)

		// Check URL
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
