package main

import (
	"context"
	"net/http"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func initiateTestMetrics() memstorage.NewMemStorageParams {
	var test_data memstorage.NewMemStorageParams

	test_data.CounterMetrics = map[string]repository.Counter{
		"test1": 1,
		"test2": 2,
		"test3": 3,
	}

	test_data.GaugeMetrics = map[string]repository.Gauge{
		"test123": 1.23,
		"test234": 2.34,
	}

	return test_data
}

func main() {

	// Step 1. Initiate Repo storage

	//var test_data = initiateTestMetrics()

	var repo, _ = memstorage.NewMemStorage(context.TODO(), nil)

	// Step 2. Get Config

	// Step 3. Start Server

	r := router.StartServer(repo)

	err := http.ListenAndServe(`:8080`, r)

	if err != nil {
		panic(err)
	}

}
