package main

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/neatflowcv/vesta/internal/app/flow"
	"github.com/neatflowcv/vesta/internal/pkg/repository/virtualbox"
)

const (
	timeout = 10 * time.Second
)

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func main() {
	log.Println("version", version())

	repository := virtualbox.NewRepository()
	service := flow.NewService(repository)
	handler := NewHandler(service)

	server := &http.Server{ //nolint:exhaustruct
		ReadHeaderTimeout: timeout,
		Addr:              ":8080",
		Handler:           handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
