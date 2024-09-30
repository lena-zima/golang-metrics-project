package serverconfig

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type ServerConfig struct {
	ServerAddr string `env:"ADDRESS"`
}

func GetConfig() (*ServerConfig, error) {

	var conf ServerConfig

	//os.Setenv("ADDRESS", "localhost:12345")

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

	return &conf, err
}

func getEnvs() (*ServerConfig, error) {

	var envs ServerConfig

	err := env.Parse(&envs)

	if err != nil {
		log.Printf("error while parsing envs: %e", err)
	}

	return &envs, err
}

func getFlags() (*ServerConfig, error) {

	var flags ServerConfig

	flag.StringVar(&flags.ServerAddr, "a", "localhost:8080", "server endpoint address")
	flag.Parse()

	return &flags, nil

}
