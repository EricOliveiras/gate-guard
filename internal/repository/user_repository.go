package repository

import (
	"context"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (model.User, error)
	Update(ctx context.Context, id uuid.UUID, updateUser request.CreateUser) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(ctx context.Context, user model.User) error {
	query := `
		INSERT INTO users 
			(id, first_name, last_name, document_id, created_at, updated_at) 
		VALUES 
			(:id, :first_name, :last_name, :document_id, :created_at, :updated_at)
		`

	_, err := ur.DB.NamedExecContext(ctx, query, &user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User

	query := "SELECT * FROM users WHERE id = ?"
	err := ur.DB.GetContext(ctx, &user, query, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, id uuid.UUID, updateUser request.CreateUser) error {
	query := "UPDATE users SET first_name = ?, last_name = ?, document_id = ? WHERE id = ?"
	_, err := ur.DB.ExecContext(ctx, query, updateUser.FirstName, updateUser.LastName, updateUser.DocumentId, id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := ur.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
