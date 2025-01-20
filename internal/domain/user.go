package domain

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(id int64) (*User, error)
	Update(user *User) (*User, error)
	Delete(id int64) error
}

type UserUseCase interface {
	Create(name, email string) (*User, error)
	GetByID(id int64) (*User, error)
	Update(id int64, name, email string) (*User, error)
	Delete(id int64) error
}
