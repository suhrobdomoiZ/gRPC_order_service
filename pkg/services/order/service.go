package order

import (
	"context"
	"errors"
	pb "homework/pkg/api/proto"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// короче, суть в том, что мне надо глянуть в прото файл, реализовать все методы которые там написаны(там буквально представлен интерфейс моего сервиса)
type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer                      //это структура с методами моего сервиса, но они все возвращают ошибку, типо мы не прописаны чел пропиши нас
	mu                                 *sync.Mutex          //мютекс для синхронизации доступа к заказам
	orders                             map[string]*pb.Order //сама мапа заказов
}

func NewOrderServiceServer() *OrderServiceServer {
	return &OrderServiceServer{
		orders: make(map[string]*pb.Order),
	}
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "order.CreateOrder: request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "order.CreateOrder: deadline exceeded")
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Item == "" {
		return nil, status.Errorf(codes.InvalidArgument, "order.CreateOrder: invalid item, must be not empty")
	}
	if req.Quantity < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "order.CreateOrder: invalid quantity, must be >0, got: %d", req.Quantity)
	}

	orderId := uuid.New().String()
	s.orders[orderId] = &pb.Order{
		Id:       orderId,
		Item:     req.Item,
		Quantity: req.Quantity,
	}
	return &pb.CreateOrderResponse{Id: orderId}, nil

}

func (s *OrderServiceServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "order.GetOrder: request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "order.GetOrder: deadline exceeded")
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "order.GetOrder: order not found")
	}
	return &pb.GetOrderResponse{Order: order}, nil

}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "order.UpdateOrder: request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "order.UpdateOrder: deadline exceeded")
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[req.Id]

	if !ok {
		return nil, status.Error(codes.NotFound, "order.UpdateOrder: order not found")
	}

	order.Quantity = req.Quantity
	order.Item = req.Item

	return &pb.UpdateOrderResponse{Order: order}, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "order.DeleteOrder: request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "order.DeleteOrder: deadline exceeded")
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.orders[req.Id]
	if !ok {
		return &pb.DeleteOrderResponse{Success: false}, status.Error(codes.NotFound, "order.DeleteOrder: order not found")
	}
	delete(s.orders, req.Id)

	return &pb.DeleteOrderResponse{Success: true}, nil

}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "order.ListOrders: request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "order.ListOrders: deadline exceeded")
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	ordersList := make([]*pb.Order, 0, len(s.orders))
	for _, order := range s.orders {
		ordersList = append(ordersList, order)
	}
	return &pb.ListOrdersResponse{Orders: ordersList}, nil
}
