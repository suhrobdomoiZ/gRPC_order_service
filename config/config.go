package config

import (
	"homework/pkg/load_config"
)

const GRPC_PORT load_config.ConfigKey = "GRPC_PORT"

type Config struct {
	grpcPort string
}

func NewConfig() *Config {
	grpcPort := GRPC_PORT.MustGet()
	return &Config{grpcPort: grpcPort}
}

func (c *Config) GRPCPort() string {
	return c.grpcPort
}
