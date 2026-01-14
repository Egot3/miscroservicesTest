package internal_test

import (
	"context"
	"testing"
	"time"

	pb "github.com/Egot3/microservicesTest/backend/apiContracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestProductService_CreateAndGet(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createResp, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:  "Test Laptop",
		Price: 1299.99,
		Stock: 10,
	})
	if err != nil {
		t.Errorf("failed to createAProduct: %v", err)
	}

	getResp, err := client.GetProduct(ctx, &pb.GetProductRequest{
		ProductId: createResp.ProductId,
	})
	if err != nil {
		t.Errorf("GetProduct failed: %v", err)
	}

	if getResp.Product.Name != "Test Laptop" {
		t.Errorf("Expected 'Test Laptop', got %s", getResp.Product.Name)
	}
}
