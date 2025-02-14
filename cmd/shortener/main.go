package main

import (
	gserver "github.com/nglmq/ozon-test/internal/app/grpc/server"
	"github.com/nglmq/ozon-test/internal/app/http/server"
	"github.com/nglmq/ozon-test/internal/app/service"
	"github.com/nglmq/ozon-test/internal/config"
	"github.com/nglmq/ozon-test/internal/storage/inmemory"
	"github.com/nglmq/ozon-test/internal/storage/postgres"
	"log"
)

func main() {
	config.ParseFlags()

	var storage service.URLRepository
	var urlService *service.URLService

	storage = inmemory.NewInMemoryURLStorage()
	urlService = service.NewURLService(storage)

	if config.InMemoryStorage && config.DBConnection == "" {
		log.Println("Using in-memory storage")

		storage = inmemory.NewInMemoryURLStorage()
		urlService = service.NewURLService(storage)
	}

	if config.DBConnection != "" {
		log.Println("Using PostgreSQL storage")

		storage = postgres.NewPostgresURLStorage(config.DBConnection)
		urlService = service.NewURLService(storage)
	}

	if config.GRPCServer {
		log.Println("Starting gRPC server")
		gserver.StartGRPCServer(urlService)
	} else {
		log.Println("Starting HTTP server")
		server.StartHTTPServer(urlService)
	}
}
