package service

import (
	"context"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/builder"
	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/repository"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
)

type UserServiceWrapper interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByDocumentID(ctx context.Context, document_id string) (*model.User, error)
	Update(ctx context.Context, id uuid.UUID, updateUser *request.CreateUser) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{Repository: repository}
}

func (us *UserService) Create(ctx context.Context, user *request.CreateUser) error {
	createUser := builder.NewUserBuilder().
		SetID(uuid.New()).
		SetFirstname(user.FirstName).
		SetLastname(user.LastName).
		SetDocumentId(user.DocumentId).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	err := us.Repository.Create(ctx, &createUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user *model.User

	user, err := us.Repository.GetByID(ctx, id)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (us *UserService) GetByDocumentID(ctx context.Context, document_id string) (*model.User, error) {
	var user *model.User

	user, err := us.Repository.GetByDocumentID(ctx, document_id)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (us *UserService) Update(ctx context.Context, id uuid.UUID, updateUser *request.CreateUser) error {
	existingUser, err := us.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	existingUser.FirstName = updateUser.FirstName
	existingUser.LastName = updateUser.LastName
	existingUser.DocumentId = updateUser.DocumentId

	err = us.Repository.Update(ctx, id, updateUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	user, err := us.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = us.Repository.Delete(ctx, user.ID)
	if err != nil {
		return err
	}

	return nil
}
