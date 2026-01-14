package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/Egot3/microservicesTest/backend/apiContracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Couldn't connect to gRPC: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		getStatus, err := client.HealthCheck(ctx, &emptypb.Empty{})
		if err == nil && getStatus.Alive {
			fmt.Printf("Server's health check succeeded (att %d/5)\n", i+1)
			os.Exit(0)
		}
		fmt.Printf("Server's health check failed (att %d/5): %v\n", i+1, err)
	}

	fmt.Println("Health check failed after 5 attempts")
	os.Exit(1)
}
