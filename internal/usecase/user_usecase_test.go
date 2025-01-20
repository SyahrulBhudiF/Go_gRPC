package usecase

import (
	"errors"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockUserRepository
	useCase  domain.UserUseCase
}

func (s *UserUseCaseTestSuite) SetupTest() {
	s.mockRepo = mocks.NewMockUserRepository(s.T())
	s.useCase = NewUserUseCase(s.mockRepo)
}

func (s *UserUseCaseTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
	s.mockRepo = mocks.NewMockUserRepository(s.T())
	s.useCase = NewUserUseCase(s.mockRepo)
}

func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (s *UserUseCaseTestSuite) TestCreate() {
	testCases := []struct {
		name          string
		input         struct{ name, email string }
		mockBehavior  func(mockRepo *mocks.MockUserRepository)
		expectedError error
	}{
		{
			name: "Success",
			input: struct{ name, email string }{
				name:  "John Doe",
				email: "john@example.com",
			},
			mockBehavior: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.EXPECT().
					Create(mock.MatchedBy(func(u *domain.User) bool {
						return u.Name == "John Doe" && u.Email == "john@example.com"
					})).
					Return(&domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Repository Error",
			input: struct{ name, email string }{
				name:  "John Doe",
				email: "john@example.com",
			},
			mockBehavior: func(mockRepo *mocks.MockUserRepository) {
				mockRepo.EXPECT().
					Create(mock.MatchedBy(func(u *domain.User) bool {
						return u.Name == "John Doe" && u.Email == "john@example.com"
					})).
					Return(nil, errors.New("repository error"))
			},
			expectedError: errors.New("repository error"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			tc.mockBehavior(s.mockRepo)

			result, err := s.useCase.Create(tc.input.name, tc.input.email)

			if tc.expectedError != nil {
				s.Equal(tc.expectedError, err)
				s.Nil(result)
			} else {
				s.NoError(err)
				s.NotNil(result)
				s.Equal(tc.input.name, result.Name)
				s.Equal(tc.input.email, result.Email)
			}
		})
	}
}
