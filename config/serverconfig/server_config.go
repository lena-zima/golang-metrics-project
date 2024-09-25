package serverconfig

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
)

type ServerConfig struct {
	Repo     repository.Repository
	ServAddr string
}

func GetConfig(serv string) (*ServerConfig, error) {

	var conf ServerConfig

	var repo, err = memstorage.NewMemStorage()

	conf.Repo = repo
	conf.ServAddr = serv

	if err != nil {
		log.Printf("err while repo creation: %e", err)
		return nil, err
	}

	return &conf, nil
}
