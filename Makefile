EXECUTABLE_NAME=order-service

MAIN_PATH=./cmd/main.go

PROTO_PATH=./internal/api/proto/order.proto

PROTO_PKG=./internal/api/proto


all: generate build

generate:
	@echo "Generating order.pb.go, order_grpc.go and order.pb.gw.go..."
	protoc -I . \
	-I ./googleapis \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
	$(PROTO_PATH)/order.proto

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
