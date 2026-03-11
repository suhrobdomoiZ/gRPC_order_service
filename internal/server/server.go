package server

import (
	"context"
	"homework/config"
	"homework/internal/services/order"
	"homework/pkg/closer"
	"homework/pkg/load_config"
	"homework/pkg/logger"
	"log"
	"log/slog"
)

type Server struct {
	logger *slog.Logger
	closer closer.Closer
	ctx    context.Context
	config config.Config

	service order.OrderServiceServer
}

//"./config/.env"
func NewServer(configPath string) *Server { //TODO: добавить closer

	ctx := context.Background()
	err := load_config.LoadDotEnv(configPath)
	if err != nil {
		log.Fatalf("server.NewServer: failed to load .env file: %v", err)
	}

	appConfig := config.NewConfig()
	logger.Setup(appConfig.EnvType())

	lg := logger.With("service_name", "order-service")

	orderServiceServer := order.NewOrderServiceServer()

	return &Server{
		logger:  lg,
		ctx:     ctx,
		config:  *appConfig,
		service: *orderServiceServer,
	}
}
