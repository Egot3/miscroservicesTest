package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Egot3/microservicesTest/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	products map[string]*pb.Product
}

func NewProductServer() *ProductServer {
	return &ProductServer{
		products: make(map[string]*pb.Product),
	}
}

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	log.Printf("GetProduct request with Id: %s", req.ProductId)

	product, exists := s.products[req.ProductId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "product was not found Id: %s", req.ProductId)
	}

	return &pb.GetProductResponse{Product: product}, nil
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Printf("Creating product: %s", req.Name)

	productId := fmt.Sprintf("prod_%d", time.Now().UnixNano())

	product := &pb.Product{
		Id:          productId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int32(req.Stock),
	}

	s.products[productId] = product

	return &pb.CreateProductResponse{ProductId: productId}, nil
}
