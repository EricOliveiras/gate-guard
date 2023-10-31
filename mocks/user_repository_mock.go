// No arquivo "user_repository_mock.go" dentro do pacote "mocks":

package mocks

import (
	"context"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByDocumentID(ctx context.Context, documentID string) (*model.User, error) {
	args := m.Called(ctx, documentID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(ctx context.Context, id uuid.UUID, user *request.CreateUser) error {
	args := m.Called(ctx, id, user)

	if err := args.Error(0); err != nil {
		return err
	}

	return nil
}

func (m *UserRepositoryMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
