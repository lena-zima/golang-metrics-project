package serverconfig

import (
	"log"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
)

type ServerConfig struct {
	Repo repository.Repository
}

func GetConfig() (*ServerConfig, error) {

	var conf ServerConfig

	var repo, err = memstorage.NewMemStorage()

	conf.Repo = repo

	if err != nil {
		log.Printf("err while repo creation: ", err)
		return nil, err
	}

	return &conf, nil
}
