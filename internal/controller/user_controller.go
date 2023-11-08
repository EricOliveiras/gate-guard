package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ericoliveiras/gate-guard/internal/request"
	"github.com/ericoliveiras/gate-guard/internal/service"
	"github.com/google/uuid"
)

type UserController struct {
	Service *service.UserService
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser request.CreateUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err := uc.Service.Create(r.Context(), &newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

id, err := uuid.Parse(userID)
if err != nil {
    http.Error(w, "Invalid user ID", http.StatusBadRequest)
    return
}

user, err := uc.Service.GetByID(r.Context(), id)
if err != nil {
    http.Error(w, "User not found", http.StatusNotFound)
    return
}

userJSON, err := json.Marshal(user)
if err != nil {
    http.Error(w, "Error encoding user to JSON", http.StatusInternalServerError)
    return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)

w.Write(userJSON)

}
