package config

import (
	"homework/pkg/load_config"
)

const (
	GRPC_PORT load_config.ConfigKey = "GRPC_PORT"
	ENV_TYPE  load_config.ConfigKey = "ENV_TYPE"
	HTTP_PORT load_config.ConfigKey = "HTTP_PORT"
)

type Config struct {
	grpcPort string
	envType  string
	httpPort string
}

func NewConfig() *Config {
	grpcPort := GRPC_PORT.MustGet()
	envType := ENV_TYPE.MustGet()
	httpPort := HTTP_PORT.MustGet()
	return &Config{
		grpcPort: grpcPort,
		envType:  envType,
		httpPort: httpPort,
	}
}

func (c *Config) GRPCPort() string {
	return c.grpcPort
}

func (c *Config) EnvType() string {
	return c.envType
}

func (c *Config) HTTPPort() string {
	return c.httpPort
}
