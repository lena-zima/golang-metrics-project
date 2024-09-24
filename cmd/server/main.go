package main

import (
	"context"
	"net/http"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func initiateTestMetrics() memstorage.NewMemStorageParams {
	var testData memstorage.NewMemStorageParams

	testData.CounterMetrics = map[string]repository.Counter{}

	testData.GaugeMetrics = map[string]repository.Gauge{}

	return testData
}

func main() {

	// Step 1. Initiate Repo storage

	var testData = initiateTestMetrics()

	var repo, _ = memstorage.NewMemStorage(context.TODO(), &testData)

	// Step 2. Get Config
	// TODO

	// Step 3. Start Server

	r := router.StartServer(repo)

	err := http.ListenAndServe(`:8080`, r)

	if err != nil {
		panic(err)
	}

}
