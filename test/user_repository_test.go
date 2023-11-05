package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/repository"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	t.Run("should create an user in database", func(t *testing.T) {
		repository, err := setupUserRepositoryTest(t)
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

		defer clearTable(t, repository.DB)

		assert.NoError(t, err)
	})

	t.Run("should return an user from database by id", func(t *testing.T) {
		repository, err := setupUserRepositoryTest(t)
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

		defer clearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, getUser.ID)
		assert.Equal(t, user.FirstName, getUser.FirstName)
		assert.Equal(t, user.LastName, getUser.LastName)
		assert.Equal(t, user.DocumentId, getUser.DocumentId)
	})

	t.Run("should return an user from database by document_id", func(t *testing.T) {
		repository, err := setupUserRepositoryTest(t)
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

		defer clearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, getUser.ID)
		assert.Equal(t, user.FirstName, getUser.FirstName)
		assert.Equal(t, user.LastName, getUser.LastName)
		assert.Equal(t, user.DocumentId, getUser.DocumentId)
	})

	t.Run("should update an user from database", func(t *testing.T) {
		repository, err := setupUserRepositoryTest(t)
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
			FirstName: "Test",
			LastName: "Test",
			DocumentId: "654321",
		}

		updatingUser, err := repository.Update(ctx, user.ID, &updateUser)
		
		defer clearTable(t, repository.DB)

		assert.NoError(t, err)
		assert.Equal(t, updateUser.FirstName, updatingUser.FirstName)
		assert.Equal(t, updateUser.LastName, updatingUser.LastName)
		assert.Equal(t, updateUser.DocumentId, updatingUser.DocumentId)
	})

	t.Run("should delete an user from database", func(t *testing.T) {
		repository, err := setupUserRepositoryTest(t)
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
		
		defer clearTable(t, repository.DB)

		assert.NoError(t, err)
	})
}

func setupUserRepositoryTest(t *testing.T) (*repository.UserRepository, error) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		return nil, err
	}

	password := os.Getenv("DB_TEST_PASS")
	driver := os.Getenv("DB_TEST_DRIVER")
	user := os.Getenv("DB_TEST_USER")
	name := os.Getenv("DB_TEST_NAME")
	host := os.Getenv("DB_TEST_HOST")
	port := os.Getenv("DB_TEST_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		t.Fatalf("Failed to connect to the database. %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	return userRepository, nil
}

func clearTable(t *testing.T, db *sqlx.DB) {
	query := "DELETE FROM users"
	_, err := db.Exec(query)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
