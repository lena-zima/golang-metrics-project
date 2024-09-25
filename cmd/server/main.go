package main

import (
	"net/http"

	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func main() {

	conf := serverconfig.GetConfig()

	r := router.StartServer(conf)

	err := http.ListenAndServe(`:8080`, r)

	if err != nil {
		panic(err)
	}

}
