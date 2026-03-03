package main

import (
	"homework/config"
	pb "homework/internal/api/proto"
	"homework/internal/middleware"
	"homework/internal/services/order"
	"homework/pkg/load_config"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	err := load_config.LoadDotEnv("./config/.env")
	if err != nil {
		log.Fatalf("main: failed to load .env file: %v", err)
	}

	appConfig := config.NewConfig()

	orderServiceServer := order.NewOrderServiceServer()
	lis, err := net.Listen("tcp", ":"+appConfig.GRPCPort())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggerInterceptor()))
	pb.RegisterOrderServiceServer(grpcServer, orderServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
