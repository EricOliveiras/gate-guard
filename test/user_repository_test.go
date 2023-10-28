package test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/ericoliveiras/gate-guard/internal/repository"
	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	t.Run("should add a new user in database", func(t *testing.T) {
		repository, mock, err := setupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		defer repository.DB.Close()

		user := model.User{
			ID:         uuid.New(),
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		mock.ExpectExec("INSERT INTO users").
			WithArgs(user.ID, user.FirstName, user.LastName, user.DocumentId, user.CreatedAt, user.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		ctx, cancel := setupTestWithContext(t, 5*time.Second)
		defer cancel()

		err = repository.Create(ctx, user)

		assert.NoError(t, err)
	})

	t.Run("should return an user from database", func(t *testing.T) {
		repository, mock, err := setupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		defer repository.DB.Close()

		id := uuid.New()

		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "document_id"}).
			AddRow(id, "Test", "Test", "123456")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id = ?")).
			WithArgs(id).
			WillReturnRows(rows)

		ctx, cancel := setupTestWithContext(t, 5*time.Second)
		defer cancel()

		user, err := repository.GetByID(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, id, user.ID)
		assert.Equal(t, "Test", user.FirstName)
		assert.Equal(t, "Test", user.LastName)
		assert.Equal(t, "123456", user.DocumentId)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})

	t.Run("should update an user", func(t *testing.T) {
		repository, mock, err := setupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		defer repository.DB.Close()

		id := uuid.New()

		sqlmock.NewRows([]string{"id", "first_name", "last_name", "document_id"}).
			AddRow(id, "Test", "Test", "123456")

		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET first_name = ?, last_name = ?, document_id = ? WHERE id = ?")).
			WithArgs("UpdateTest", "UpdateTest", "654321", id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		updateUser := request.CreateUser{
			FirstName:  "UpdateTest",
			LastName:   "UpdateTest",
			DocumentId: "654321",
		}

		ctx, cancel := setupTestWithContext(t, 5*time.Second)
		defer cancel()

		err = repository.Update(ctx, id, updateUser)

		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})

	t.Run("should delete an user from database", func(t *testing.T) {
		repository, mock, err := setupUserRepositoryTest(t)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		defer repository.DB.Close()

		id := uuid.New()

		sqlmock.NewRows([]string{"id", "first_name", "last_name", "document_id"}).
			AddRow(id, "Test", "Test", "123456")

		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = ?")).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := setupTestWithContext(t, 5*time.Second)
		defer cancel()

		err = repository.Delete(ctx, id)

		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})
}

func setupUserRepositoryTest(t *testing.T) (userRepository *repository.UserRepository, mock sqlmock.Sqlmock, err error) {
	testDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating database mock: %v", err)
	}

	sqlxDB := sqlx.NewDb(testDB, "sqlmock_db")
	if err != nil {
		t.Fatalf("Error creating database: %v", err)
	}

	repository := repository.NewUserRepository(sqlxDB)

	return repository, mock, nil
}

func setupTestWithContext(t *testing.T, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx, cancel
}
