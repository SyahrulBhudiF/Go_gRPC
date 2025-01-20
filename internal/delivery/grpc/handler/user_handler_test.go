package handler

import (
	"context"
	"errors"
	pb "github.com/SyahrulBhudiF/Go_gRPC/internal/delivery/grpc/proto"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserHandlerTestSuite struct {
	suite.Suite
	mockUseCase *mocks.MockUserUseCase
	handler     *UserHandler
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.mockUseCase = mocks.NewMockUserUseCase(s.T())
	s.handler = NewUserHandler(s.mockUseCase)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) TestCreateUser() {
	testCases := []struct {
		name          string
		request       *pb.CreateUserRequest
		mockBehavior  func()
		expectedError error
	}{
		{
			name: "Success",
			request: &pb.CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockBehavior: func() {
				s.mockUseCase.EXPECT().
					Create("John Doe", "john@example.com").
					Return(&domain.User{
						ID:    1,
						Name:  "John Doe",
						Email: "john@example.com",
					}, nil)
			},
			expectedError: nil,
		},
		{
			name: "UseCase Error",
			request: &pb.CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockBehavior: func() {
				// Reset mock state before setting up the behavior again
				s.mockUseCase.EXPECT().
					Create("John Doe", "john@example.com").
					Return(nil, errors.New("usecase error"))
			},
			expectedError: errors.New("usecase error"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			tc.mockBehavior()

			resp, err := s.handler.CreateUser(context.Background(), tc.request)

			if tc.expectedError != nil {
				s.Equal(tc.expectedError.Error(), err.Error())
				s.Nil(resp)
			} else {
				s.NoError(err)
				s.NotNil(resp)
				s.Equal(tc.request.Name, resp.User.Name)
				s.Equal(tc.request.Email, resp.User.Email)
			}
		})
	}
}
