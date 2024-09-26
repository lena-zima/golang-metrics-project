package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/lena-zima/golang-metrics-project/config/agentconfig"
	"github.com/lena-zima/golang-metrics-project/internal/agent"
)

func main() {

	srvAddr := flag.String("a", "localhost:8080", "server endpoint address")
	repInt := flag.Int("r", 10, "report interval")
	pollInt := flag.Int("p", 2, "poll interval")

	flag.Parse()

	srvEnv, srvEnvExists := os.LookupEnv("ADDRESS")

	if srvEnvExists == true {
		*srvAddr = srvEnv
	}

	repEnv, repEnvExists := os.LookupEnv("ADDRESS")

	if repEnvExists == true {
		var err error
		*repInt, err = strconv.Atoi(repEnv)

		if err != nil {
			log.Printf("failed to convert env variable %e", err)
		}
	}

	pollEnv, pollEnvExists := os.LookupEnv("ADDRESS")

	if pollEnvExists == true {
		var err error
		*pollInt, err = strconv.Atoi(pollEnv)

		if err != nil {
			log.Printf("failed to convert env variable %e", err)
		}
	}

	conf, err := agentconfig.GetConfig("http://"+*srvAddr, *repInt, *pollInt)

	if err != nil {
		log.Fatalf("failed to get agent config %e", err)
	}

	a, err := agent.NewAgent(conf)

	if err != nil {
		log.Fatalf("failed to create agent %e", err)
	}

	a.RunJob()

}
