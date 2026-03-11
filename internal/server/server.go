package server

import (
	"context"
	"homework/config"
	pb "homework/internal/api/proto"
	"homework/internal/middleware"
	"homework/internal/services/order"
	"homework/pkg/closer"
	"homework/pkg/load_config"
	"homework/pkg/logger"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	logger *slog.Logger
	closer closer.Closer
	ctx    context.Context
	config config.Config

	service order.OrderServiceServer
}

// "./config/.env"
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

func (s *Server) Run() error { //TODO:вернуться после closer

	errCh := make(chan error, 1)

	lis, err := net.Listen("tcp", ":"+s.config.GRPCPort())

	if err != nil {
		s.logger.Error("server.Run: failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggerInterceptor(*s.logger)))
	pb.RegisterOrderServiceServer(grpcServer, &s.service)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
