package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GrpcUrl string `envconfig:"GRPC_URL" default:"localhost:50051"`
	RestUrl string `envconfig:"REST_URL" default:"localhost:50052"`
}

func New() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("func envconfig.Process; cfg %v;", cfg)
	}
	return cfg
}
