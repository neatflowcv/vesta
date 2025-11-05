package main

import (
	"flag"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/neatflowcv/vesta/internal/app/flow"
	"github.com/neatflowcv/vesta/internal/pkg/client/virtualbox"
	"github.com/neatflowcv/vesta/internal/pkg/repository/redis"
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

	// Redis flags
	redisAddr := flag.String("redis-addr", "127.0.0.1:6379", "Redis address host:port")
	redisPassword := flag.String("redis-password", "", "Redis password")
	redisDB := flag.Int("redis-db", 0, "Redis DB index")

	flag.Parse()

	// Initialize Redis repository (currently not wired into handlers; prepared for future use)
	repo, err := redis.NewRepository(*redisAddr, *redisPassword, *redisDB)
	if err != nil {
		log.Fatal(err)
	}

	client := virtualbox.NewClient()
	service := flow.NewService(client, repo)
	handler := NewHandler(service)

	server := &http.Server{ //nolint:exhaustruct
		ReadHeaderTimeout: timeout,
		Addr:              ":8080",
		Handler:           handler,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
