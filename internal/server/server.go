package server

import (
	"context"
	"errors"
	"fmt"
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
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	logger *slog.Logger
	closer *closer.Closer
	ctx    context.Context
	config config.Config

	service order.OrderServiceServer
}

// "./config/.env"
func NewServer(configPath string) *Server {

	ctx := context.Background()
	err := load_config.LoadDotEnv(configPath)
	if err != nil {
		log.Fatalf("server.NewServer: failed to load .env file: %v", err)
	}

	appConfig := config.NewConfig()
	logger.Setup(appConfig.EnvType())

	lg := logger.With("service_name", "order-service")

	orderServiceServer := order.NewOrderServiceServer()

	serverCloser := closer.New(*lg)

	return &Server{
		logger:  lg,
		ctx:     ctx,
		config:  *appConfig,
		service: *orderServiceServer,
		closer:  serverCloser,
	}
}

func (s *Server) Run() error {

	lis, err := net.Listen("tcp", ":"+s.config.GRPCPort())
	if err != nil {
		s.logger.Error("server.Run: failed to listen: %v", err)
		return fmt.Errorf("server.Run: %v", err)
	}

	s.closer.AddFunc("grpc listener", func() {
		_ = lis.Close()
	})

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggerInterceptor(*s.logger)))
	pb.RegisterOrderServiceServer(grpcServer, &s.service)

	s.closer.Add("grpc server", func(ctx context.Context) error {
		done := make(chan struct{})

		go func() {
			grpcServer.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			return nil
		case <-ctx.Done():
			grpcServer.Stop()
			<-done
			return ctx.Err()
		}
	})
	errCh := make(chan error, 1)

	go func() {
		if err := grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			errCh <- fmt.Errorf("server.Run: %v", err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return err
	case signal := <-sigCh:
		s.logger.Info("server.Run: received signal", "signal", signal.String())
		shutdownctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.closer.Close(shutdownctx); err != nil && errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("server.Run: gracefull shutdown  %v", err)
		}
	}
	s.logger.Info("server stopped")

	return nil
}
