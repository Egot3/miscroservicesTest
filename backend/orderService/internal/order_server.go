package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Egot3/microservicesTest/backend/apiContracts"
)

type OrderServer struct {
	pb.UnimplementedProductServiceServer
	productClient pb.ProductServiceClient
	orders        map[string]*pb.Order
}

func NewProductServer() *OrderServer {
	return &OrderServer{
		orders: make(map[string]*pb.Order),
	}
}

func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	log.Printf("GetOrder request with Id: %v", req.OrderId)

	order, exists := s.orders[req.OrderId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "product was not found Id: %s", req.OrderId)
	}

	return &pb.GetOrderResponse{Order: order}, nil
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("Creating order for: %s", req.CustomerId)

	orderId := fmt.Sprintf("order_%d", time.Now().UnixNano())
	subtotal, discounts, shipping_fee, tax, grand_total := 0, 0, 0, 0, 0
	for _, item := range req.Items {

	}

	//order := &pb.Order{
	//	Id:
	//} make with bridge
}
