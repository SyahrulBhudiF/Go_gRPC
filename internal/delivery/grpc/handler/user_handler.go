package handler

import (
	"context"
	pb "github.com/SyahrulBhudiF/Go_gRPC/internal/delivery/grpc/proto"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userUseCase domain.UserUseCase
}

func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.userUseCase.Create(req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userUseCase.GetByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := h.userUseCase.Update(req.Id, req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.userUseCase.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
