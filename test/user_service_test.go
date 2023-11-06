package test

import (
	"context"
	"testing"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/ericoliveiras/gate-guard/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	t.Run("should create an user", func(t *testing.T) {
		service, err := SetupUserServiceTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer service.Repository.DB.Close()

		ctx := context.Background()

		user := request.CreateUser{
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
		}

		err = service.Create(ctx, &user)

		defer ClearTable(t, service.Repository.DB)

		assert.NoError(t, err)
	})

	t.Run("should return an user by id", func(t *testing.T) {
		service, err := SetupUserServiceTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer service.Repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = service.Repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		getUser, err := service.GetByID(ctx, user.ID)

		defer ClearTable(t, service.Repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, getUser.FirstName, user.FirstName)
		assert.Equal(t, getUser.LastName, user.LastName)
		assert.Equal(t, getUser.DocumentId, user.DocumentId)
	})

	t.Run("should return an user by document_id", func(t *testing.T) {
		service, err := SetupUserServiceTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer service.Repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = service.Repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		getUser, err := service.GetByDocumentID(ctx, user.DocumentId)

		defer ClearTable(t, service.Repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, getUser.FirstName, user.FirstName)
		assert.Equal(t, getUser.LastName, user.LastName)
		assert.Equal(t, getUser.DocumentId, user.DocumentId)
	})

	t.Run("should update an user", func(t *testing.T) {
		service, err := SetupUserServiceTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer service.Repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = service.Repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		updateUser := request.CreateUser{
			FirstName:  "UpdateTestFirstName",
			LastName:   "UpdateTestLastName",
			DocumentId: "654321",
		}

		updatedUser, err := service.Update(ctx, user.ID, &updateUser)

		defer ClearTable(t, service.Repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, updatedUser.FirstName, updateUser.FirstName)
		assert.Equal(t, updatedUser.LastName, updateUser.LastName)
		assert.Equal(t, updatedUser.DocumentId, updateUser.DocumentId)
		assert.NotEqual(t, updatedUser.UpdatedAt, user.UpdatedAt)
	})

	t.Run("should delete an user", func(t *testing.T) {
		service, err := SetupUserServiceTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer service.Repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = service.Repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		err = service.Delete(ctx, user.ID)

		defer ClearTable(t, service.Repository.DB)

		assert.NoError(t, err)
	})
}

func SetupUserServiceTest(t *testing.T) (*service.UserService, error) {
	repository, err := SetupUserRepositoryTest(t)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}

	service := service.NewUserService(repository)

	return service, nil
}
