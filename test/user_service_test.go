package test

import (
	"context"
	"testing"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/ericoliveiras/gate-guard/internal/service"
	"github.com/ericoliveiras/gate-guard/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	t.Run("should create an user", func(t *testing.T) {
		userRepositoryMock, userService := setupUserServiceMock(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		userRepositoryMock.On("Create", ctx, mock.Anything).Return(nil)

		user := &request.CreateUser{
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
		}

		err := userService.Create(ctx, user)

		assert.NoError(t, err)

		userRepositoryMock.AssertExpectations(t)
	})

	t.Run("should get an user by id", func(t *testing.T) {
		userRepositoryMock, userService := setupUserServiceMock(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := uuid.New()

		userRepositoryMock.On("GetByID", ctx, id).Return(&model.User{
			ID:         id,
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)

		user, err := userService.GetByID(ctx, id)

		assert.NoError(t, err)

		userRepositoryMock.AssertExpectations(t)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, "Test", user.FirstName)
		assert.Equal(t, "Test", user.LastName)
		assert.Equal(t, "123456", user.DocumentId)
	})

	t.Run("should get an user by document_id", func(t *testing.T) {
		userRepositoryMock, userService := setupUserServiceMock(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := uuid.New()
		document_id := "123456"

		userRepositoryMock.On("GetByDocumentID", ctx, document_id).Return(&model.User{
			ID:         id,
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: document_id,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)

		user, err := userService.GetByDocumentID(ctx, document_id)

		assert.NoError(t, err)

		userRepositoryMock.AssertExpectations(t)

		assert.Equal(t, document_id, user.DocumentId)
	})

	t.Run("should update an user", func(t *testing.T) {
		userRepositoryMock, userService := setupUserServiceMock(t)

		id := uuid.New()
		updatedUser := &request.CreateUser{
			FirstName:  "UpdatedTest",
			LastName:   "UpdatedTest",
			DocumentId: "123456",
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		userRepositoryMock.On("GetByID", ctx, id).Return(&model.User{
			ID:         id,
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)

		userRepositoryMock.On("Update", ctx, id, updatedUser).Return(nil)

		err := userService.Update(ctx, id, updatedUser)

		assert.NoError(t, err)

		userRepositoryMock.AssertExpectations(t)
	})

	t.Run("should delete an user", func(t *testing.T) {
		userRepositoryMock, userService := setupUserServiceMock(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := uuid.New()

		userRepositoryMock.On("GetByID", ctx, id).Return(&model.User{
			ID:         id,
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)

		userRepositoryMock.On("Delete", ctx, id).Return(nil)

		err := userService.Delete(ctx, id)

		assert.NoError(t, err)

		userRepositoryMock.AssertExpectations(t)
	})
}

func setupUserServiceMock(t *testing.T) (*mocks.UserRepositoryMock, *service.UserService) {
	userRepositoryMock := &mocks.UserRepositoryMock{}

	userRepositoryWrapper := &mocks.UserRepositoryMockWrapper{UserRepositoryMock: userRepositoryMock}

	userService := service.NewUserService(userRepositoryWrapper)

	return userRepositoryMock, userService
}
