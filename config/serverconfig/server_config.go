package serverconfig

import (
	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/lena-zima/golang-metrics-project/internal/repository/memstorage"
)

type ServerConfig struct {
	Repo repository.Repository
}

func GetConfig() *ServerConfig {

	var conf ServerConfig

	var repo, err = memstorage.NewMemStorage()

	conf.Repo = repo

	if err != nil {
		panic("AA")
	}

	return &conf
}
