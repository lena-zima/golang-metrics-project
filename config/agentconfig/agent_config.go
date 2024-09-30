package agentconfig

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type AgentConfig struct {
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	ServerAddr     string `env:"ADDRESS"`
}

func GetConfig() (*AgentConfig, error) {

	var conf AgentConfig

	//os.Setenv("ADDRESS", "localhost:12346")

	envs, err := getEnvs()

	if err != nil {
		log.Printf("error while parsing envs: %e", err)
	}

	flags, err := getFlags()

	if err != nil {
		log.Printf("error while parsing flags: %e", err)
	}

	if envs.ServerAddr == "" {
		conf.ServerAddr = flags.ServerAddr
	} else {
		conf.ServerAddr = envs.ServerAddr
	}

	conf.ServerAddr = "http://" + conf.ServerAddr

	if envs.PollInterval == 0 {
		conf.PollInterval = flags.PollInterval
	} else {
		conf.PollInterval = envs.PollInterval
	}

	if envs.ReportInterval == 0 {
		conf.ReportInterval = flags.ReportInterval
	} else {
		conf.ReportInterval = envs.ReportInterval
	}

	return &conf, err
}

func getEnvs() (*AgentConfig, error) {

	var envs AgentConfig

	err := env.Parse(&envs)

	if err != nil {
		log.Printf("error while parsing envs: %e", err)
	}

	return &envs, err
}

func getFlags() (*AgentConfig, error) {

	var flags AgentConfig

	flag.StringVar(&flags.ServerAddr, "a", "localhost:8080", "server endpoint address")
	flag.IntVar(&flags.PollInterval, "p", 2, "poll interval")
	flag.IntVar(&flags.ReportInterval, "r", 10, "report interval")

	flag.Parse()

	return &flags, nil

}
