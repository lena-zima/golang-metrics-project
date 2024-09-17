package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/lena-zima/golang-metrics-project/global"
)

func UpdateHandler(res http.ResponseWriter, req *http.Request) {
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

				global.St.AddGauge(metricName, float64(value))

				res.WriteHeader(http.StatusOK)

			case "counter":
				value, err := strconv.Atoi(metricValue)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					return
				}

				global.St.AddCounter(metricName, int64(value))

				res.WriteHeader(http.StatusOK)

			default:
				res.WriteHeader(http.StatusBadRequest)
			}

		}
	}

}
