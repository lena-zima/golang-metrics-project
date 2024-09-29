package serverconfig

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type ServerConfig struct {
	ServerAddr string
}

type vars struct {
	srvAddr string `env:"ADDRESS"`
}

func GetConfig() (*ServerConfig, error) {

	var conf ServerConfig

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
	flag.Parse()

	return &flags, nil

}
