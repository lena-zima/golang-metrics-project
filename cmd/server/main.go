package main

import (
	"github.com/lena-zima/golang-metrics-project/config/serverconfig"
	"github.com/lena-zima/golang-metrics-project/internal/router"
)

func main() {

	conf := serverconfig.GetConfig()

	r, err := router.NewServer(conf)

	if err != nil {
		panic("AA")
	}

	router.StartServer(r)

}
