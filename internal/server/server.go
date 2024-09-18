package server

import (
	"net/http"

	"github.com/lena-zima/golang-metrics-project/internal/handlers"
)

func StartServer() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handlers.UpdateHandler)

	return mux
}
