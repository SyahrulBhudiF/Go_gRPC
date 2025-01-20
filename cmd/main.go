package main

import (
	"github.com/SyahrulBhudiF/Go_gRPC/internal/config"
	"log"
)

func main() {
	db, _, grpcServer, lis, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	defer db.Close()

	log.Printf("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
