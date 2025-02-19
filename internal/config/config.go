package config

import (
	"flag"
)

var (
	InMemoryStorage bool
	DBConnection    string
	BaseURL         string
	Port            string
	GRPCServer      bool
)

func ParseFlags() {
	flag.BoolVar(&InMemoryStorage, "m", true, "in memory storage")
	flag.StringVar(&DBConnection, "d", "", "postgres connection url")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "base url")
	flag.StringVar(&Port, "p", ":8080", "port to run server")
	flag.BoolVar(&GRPCServer, "g", false, "run gRPC server")
	flag.Parse()
}
