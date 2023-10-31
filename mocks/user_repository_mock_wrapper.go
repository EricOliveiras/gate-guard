package mocks

import (
	"context"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
)

type UserRepositoryMockWrapper struct {
	UserRepositoryMock *UserRepositoryMock
}

func (w *UserRepositoryMockWrapper) Create(ctx context.Context, user *model.User) error {
	return w.UserRepositoryMock.Create(ctx, user)
}

func (w *UserRepositoryMockWrapper) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := w.UserRepositoryMock.Called(ctx, id)

	if err := args.Error(1); err != nil {
		return nil, err
	}

	user := args.Get(0).(*model.User)
	return user, nil
}

func (w *UserRepositoryMockWrapper) GetByDocumentID(ctx context.Context, document_id string) (*model.User, error) {
	args := w.UserRepositoryMock.Called(ctx, document_id)

	if err := args.Error(1); err != nil {
		return nil, err
	}

	user := args.Get(0).(*model.User)
	return user, nil
}

func (w *UserRepositoryMockWrapper) Update(ctx context.Context, id uuid.UUID, user *request.CreateUser) error {
	return w.UserRepositoryMock.Update(ctx, id, user)
}

func (w *UserRepositoryMockWrapper) Delete(ctx context.Context, id uuid.UUID) error {
	return w.UserRepositoryMock.Delete(ctx, id)
}
