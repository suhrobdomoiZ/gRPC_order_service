package main

import (
	"fmt"
	pb "homework/pkg/api/proto"
	"homework/pkg/services/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	orderServiceServer := order.NewOrderServiceServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderServiceServer)
	fmt.Println("Server is running on :50051") //УРААААААААААААААААААААААААААААААААААААА
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
