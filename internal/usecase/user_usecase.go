package usecase

import "github.com/SyahrulBhudiF/Go_gRPC/internal/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(name, email string) (*domain.User, error) {
	user := &domain.User{
		Name:  name,
		Email: email,
	}
	return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetByID(id int64) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUseCase) Update(id int64, name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return uc.userRepo.Update(user)
}

func (uc *userUseCase) Delete(id int64) error {
	return uc.userRepo.Delete(id)
}
