package config

import (
	"homework/pkg/load_config"
)

const (
	GRPC_PORT load_config.ConfigKey = "GRPC_PORT"
	ENV_TYPE  load_config.ConfigKey = "ENV_TYPE"
	HTTP_PORT load_config.ConfigKey = "HTTP_PORT"

	DB_USERNAME load_config.ConfigKey = "DB_USERNAME"
	DB_PASSWORD load_config.ConfigKey = "DB_PASSWORD"
	DB_HOST     load_config.ConfigKey = "DB_HOST"
	DB_PORT     load_config.ConfigKey = "DB_PORT"

	DB_URL load_config.ConfigKey = "DB_URL"
)

type configDB struct {
	username string
	password string
	host     string
	port     string
	dsn      string
}

func NewConfigDB() *configDB {
	return &configDB{
		username: DB_USERNAME.MustGet(),
		password: DB_PASSWORD.MustGet(),
		host:     DB_HOST.MustGet(),
		port:     DB_PORT.MustGet(),
		dsn:      DB_URL.MustGet(),
	}
}

func (db configDB) Username() string {
	return db.username
}
func (db configDB) Password() string {
	return db.password
}
func (db configDB) Host() string {
	return db.host
}
func (db configDB) Port() string {
	return db.port
}
func (db configDB) DSN() string {
	return db.dsn
}

type Config struct {
	grpcPort string
	envType  string
	httpPort string
	db       *configDB
}

func NewConfig() *Config {
	grpcPort := GRPC_PORT.MustGet()
	envType := ENV_TYPE.MustGet()
	httpPort := HTTP_PORT.MustGet()

	return &Config{
		grpcPort: grpcPort,
		envType:  envType,
		httpPort: httpPort,
		db:       NewConfigDB(),
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

func (c *Config) DB() *configDB {
	return c.db
}
