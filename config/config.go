package config

import (
	"homework/pkg/load_config"
)

const (
	GRPC_PORT load_config.ConfigKey = "GRPC_PORT"
	ENV_TYPE  load_config.ConfigKey = "ENV_TYPE"
)

type Config struct {
	grpcPort string
	envType  string
}

func NewConfig() *Config {
	grpcPort := GRPC_PORT.MustGet()
	envType := ENV_TYPE.MustGet()
	return &Config{
		grpcPort: grpcPort,
		envType:  envType,
	}
}

func (c *Config) GRPCPort() string {
	return c.grpcPort
}

func (c *Config) EnvType() string {
	return c.envType
}
