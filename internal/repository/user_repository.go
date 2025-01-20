package repository

import (
	"database/sql"
	"github.com/SyahrulBhudiF/Go_gRPC/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *domain.User) (*domain.User, error) {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email`
	err := r.db.QueryRow(query, user.Name, user.Email, user.ID).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
