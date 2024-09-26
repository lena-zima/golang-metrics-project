package serverconfig

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
)

type ServerConfig struct {
	Repo     repository.Repository
	ServAddr string
}

func GetConfig() (*ServerConfig, error) {

	var conf ServerConfig

	var repo, err = memstorage.NewMemStorage()

	if err != nil {
		log.Printf("err while repo creation: %e", err)
		return nil, err
	}

	conf.Repo = repo

	srvAddr, err := getEnv("srvAddr")

	if err != nil {
		log.Printf("failed to get server address env: %e", err)
	}

	conf.ServAddr = srvAddr

	if srvAddr == "" {
		srvAddr, err = getFlag("srvAddr")

		conf.ServAddr = srvAddr

		if err != nil {
			log.Printf("failed to get server address flag: %e", err)
		}
	}

	return &conf, err
}

func getEnv(name string) (string, error) {

	switch name {
	case "srvAddr":
		srvEnv, srvEnvExists := os.LookupEnv("ADDRESS")
		if srvEnvExists {
			return srvEnv, nil
		}
	default:
		err := errors.New("unknown env")
		log.Printf("env parsing error: %e", err)
		return "", err
	}
	return "", nil
}

func getFlag(name string) (string, error) {
	switch name {
	case "srvAddr":
		srvAddr := flag.String("a", "localhost:8080", "server endpoint address")
		flag.Parse()
		return *srvAddr, nil
	default:
		err := errors.New("unknown flag")
		log.Printf("env parsing flags: %e", err)
		return "", err
	}

}
