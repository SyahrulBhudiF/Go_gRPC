package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	db       *sql.DB
	userRepo domain.UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	var err error
	s.db, s.mock, err = sqlmock.New()
	s.Require().NoError(err)
	s.userRepo = NewUserRepository(s.db)
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	s.db.Close()
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestCreate() {
	testCases := []struct {
		name          string
		input         *domain.User
		mockBehavior  func()
		expectedError error
	}{
		{
			name: "Success",
			input: &domain.User{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				s.mock.ExpectQuery("INSERT INTO users").
					WithArgs("John Doe", "john@example.com").
					WillReturnRows(rows)
			},
			expectedError: nil,
		},
		{
			name: "Database Error",
			input: &domain.User{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockBehavior: func() {
				s.mock.ExpectQuery("INSERT INTO users").
					WithArgs("John Doe", "john@example.com").
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.mockBehavior()

			result, err := s.userRepo.Create(tc.input)

			if tc.expectedError != nil {
				s.Equal(tc.expectedError, err)
				s.Nil(result)
			} else {
				s.NoError(err)
				s.NotNil(result)
				s.Equal(int64(1), result.ID)
			}
		})
	}
}
