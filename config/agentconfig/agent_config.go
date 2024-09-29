package agentconfig

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type AgentConfig struct {
	PollInterval   int
	ReportInterval int
	ServerAddr     string
}

type vars struct {
	pollInt int    `env:"POLL_INTERVAL"`
	repInt  int    `env:"REPORT_INTERVAL"`
	srvAddr string `env:"ADDRESS"`
}

func GetConfig() (*AgentConfig, error) {

	var conf AgentConfig

	envs, err := getEnvs()

	if err != nil {
		log.Printf("error while parsing envs: %e", err)
	}

	flags, err := getFlags()

	if err != nil {
		log.Printf("error while parsing flags: %e", err)
	}

	if envs.srvAddr == "" {
		conf.ServerAddr = flags.srvAddr
	}

	conf.ServerAddr = "http://" + conf.ServerAddr

	if envs.pollInt == 0 {
		conf.PollInterval = flags.pollInt
	}

	if envs.repInt == 0 {
		conf.ReportInterval = flags.repInt
	}

	return &conf, err
}

func getEnvs() (*vars, error) {

	var envs vars

	err := env.Parse(&envs)

	if err != nil {
		log.Printf("error while parsing envs: %e", err)
	}

	return &envs, err
}

func getFlags() (*vars, error) {

	var flags vars

	flag.StringVar(&flags.srvAddr, "a", "localhost:8080", "server endpoint address")
	flag.IntVar(&flags.pollInt, "p", 2, "poll interval")
	flag.IntVar(&flags.repInt, "r", 10, "report interval")

	flag.Parse()

	return &flags, nil

}
