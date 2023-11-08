package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ericoliveiras/gate-guard/internal/controller"
	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserController(t *testing.T) {
	t.Run("should return status 201 when a user is created", func(t *testing.T) {
		controller, db, err := SetupUserControllerTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer db.Close()

		requestBody := []byte(`{"first_name": "Test", "last_name": "Test", "document_id": "123456"}`)
		req, err := http.NewRequest("POST", "/create-user", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		ctx := req.Context()

		recorder := httptest.NewRecorder()

		controller.Create(recorder, req.WithContext(ctx))

		defer ClearTable(t, db)

		assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)
	})

	t.Run("should return status 200 and a user in response", func(t *testing.T) {
		controller, db, err := SetupUserControllerTest(t)
		if err != nil {
			t.Fatalf("Erro: %v", err)
		}
		defer db.Close()

		ctx := context.Background()

		newUser := model.User{
			ID:         uuid.New(),
			FirstName:  "Test",
			LastName:   "Test",
			DocumentId: "123456",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		controller.Service.Repository.Create(ctx, &newUser)

		url := fmt.Sprintf("/get-user?id=%v", newUser.ID)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		recorder := httptest.NewRecorder()
		controller.GetUserById(recorder, req)

		defer ClearTable(t, db)

		assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

		var responseUser model.User
		err = json.Unmarshal(recorder.Body.Bytes(), &responseUser)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		assert.Equal(t, responseUser.ID, newUser.ID)
		assert.Equal(t, responseUser.FirstName, newUser.FirstName)
	})
}

func SetupUserControllerTest(t *testing.T) (*controller.UserController, *sqlx.DB, error) {
	service, err := SetupUserServiceTest(t)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}

	controller := controller.UserController{Service: service}
	db := service.Repository.DB

	return &controller, db, nil
}
