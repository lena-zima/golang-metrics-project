package main

import (
	"net/http"

	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func main() {

	// Step 1. Initiate Repo storage

	var repo, _ = memstorage.NewMemStorage()

	// Step 2. Get Config
	// TODO

	// Step 3. Start Server

	r := router.StartServer(repo)

	err := http.ListenAndServe(`:8080`, r)

	if err != nil {
		panic(err)
	}

}
