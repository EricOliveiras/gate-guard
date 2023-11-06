package test

import (
	"context"
	"testing"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/repository"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	t.Run("should create an user in database", func(t *testing.T) {
		repository, err := SetupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = repository.Create(ctx, &user)

		defer ClearTable(t, repository.DB)

		assert.NoError(t, err)
	})

	t.Run("should return an user from database by id", func(t *testing.T) {
		repository, err := SetupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		getUser, err := repository.GetByID(ctx, user.ID)

		defer ClearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, getUser.ID)
		assert.Equal(t, user.FirstName, getUser.FirstName)
		assert.Equal(t, user.LastName, getUser.LastName)
		assert.Equal(t, user.DocumentId, getUser.DocumentId)
	})

	t.Run("should return an user from database by document_id", func(t *testing.T) {
		repository, err := SetupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		getUser, err := repository.GetByDocumentID(ctx, "123456")

		defer ClearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, getUser.ID)
		assert.Equal(t, user.FirstName, getUser.FirstName)
		assert.Equal(t, user.LastName, getUser.LastName)
		assert.Equal(t, user.DocumentId, getUser.DocumentId)
	})

	t.Run("should update an user from database", func(t *testing.T) {
		repository, err := SetupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		updateUser := request.CreateUser{
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "654321",
		}

		updatingUser, err := repository.Update(ctx, user.ID, &updateUser)

		defer ClearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, updateUser.FirstName, updatingUser.FirstName)
		assert.Equal(t, updateUser.LastName, updatingUser.LastName)
		assert.Equal(t, updateUser.DocumentId, updatingUser.DocumentId)
	})

	t.Run("should delete an user from database", func(t *testing.T) {
		repository, err := SetupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer repository.DB.Close()

		ctx := context.Background()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "TestFirstName",
			LastName:   "TestLastName",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = repository.Create(ctx, &user)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}

		err = repository.Delete(ctx, user.ID)

		defer ClearTable(t, repository.DB)

		assert.NoError(t, err)
	})
}

func SetupUserRepositoryTest(t *testing.T) (*repository.UserRepository, error) {
	driver, dsn, err := SetupDbTest()
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		t.Fatalf("Failed to connect to the database. %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	return userRepository, nil
}

func ClearTable(t *testing.T, db *sqlx.DB) {
	query := "DELETE FROM users"
	_, err := db.Exec(query)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
