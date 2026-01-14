package main

import (
	"log"
	"net"

	"github.com/Egot3/microservicesTest/backend/productService/internal"
	pb "github.com/Egot3/microservicesTest/backend/apiContracts"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	productServer := internal.NewProductServer()
	pb.RegisterProductServiceServer(grpcServer, productServer)

	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("Product Service gRPC server on %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
