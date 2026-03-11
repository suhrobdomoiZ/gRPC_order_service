package main

import (
	"homework/config"
	pb "homework/internal/api/proto"
	"homework/internal/middleware"
	"homework/internal/services/order"
	"homework/pkg/load_config"
	"homework/pkg/logger"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {

	err := load_config.LoadDotEnv("./config/.env")
	if err != nil {
		log.Fatalf("main: failed to load .env file: %v", err)
	}

	appConfig := config.NewConfig()
	logger.Setup(appConfig.EnvType())

	lg := logger.With("service_name", "order-service")

	orderServiceServer := order.NewOrderServiceServer()
	lis, err := net.Listen("tcp", ":"+appConfig.GRPCPort())

	if err != nil {
		lg.Error("main: failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggerInterceptor(*lg)))
	pb.RegisterOrderServiceServer(grpcServer, orderServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
