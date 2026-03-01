EXECUTABLE_NAME=order-service

MAIN_PATH=./cmd/main.go

PROTO_PATH=./internal/api/proto/order.proto

PROTO_PKG=./internal/api/proto


all: generate build

generate:
	@echo "Generating order.pb.go and order_grpc.go, using order.proto..."
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           $(PROTO_PATH)

build:
	@echo "Building $(EXECUTABLE_NAME)..."
	go build -o $(EXECUTABLE_NAME) $(MAIN_PATH)


run: build
	@echo "Starting server..."
	./$(EXECUTABLE_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(EXECUTABLE_NAME)

help:
	@echo "Available commands:"
	@echo "  make generate  - Generate gRPC code from proto"
	@echo "  make build     - Build the binary"
	@echo "  make run       - Build and run the server"
	@echo "  make clean     - Remove binary file"