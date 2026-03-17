package postgres

import (
	"homework/pkg/load_config"
	"time"
)

const (
	MAX_CONNS              load_config.ConfigKey = "MAX_CONNS"
	MAX_CONN_LIFE_TIME     load_config.ConfigKey = "MAX_CONN_LIFE_TIME"
	MAX_CONN_IDLE_TIME     load_config.ConfigKey = "MAX_CONN_IDLE_TIME"
	CONNECT_TIMEOUT        load_config.ConfigKey = "CONNECT_TIMEOUT"
	defaultMaxConns                              = 10
	defaultMaxConnLifeTime                       = time.Hour
	defaultMaxConnIdleTime                       = 30 * time.Minute
	defaultConnectTimeout                        = 5 * time.Second
)

type Options struct {
	MaxConns        int
	MaxConnLifeTime time.Duration
	MaxConnIdleTime time.Duration
	ConnectTimeout  time.Duration
}

func NewOptions() *Options {

}
