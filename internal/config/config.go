package config

import (
	"database/sql"
	"fmt"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/delivery/grpc/handler"
	pb "github.com/SyahrulBhudiF/Go_gRPC/internal/delivery/grpc/proto"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/repository"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/usecase"
	"google.golang.org/grpc"
	"net"
	"os"
)

func LoadConfig() (*sql.DB, *handler.UserHandler, *grpc.Server, net.Listener, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	return db, userHandler, grpcServer, lis, nil
}
