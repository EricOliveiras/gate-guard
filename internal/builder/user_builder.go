package builder

import (
	"time"

	"github.com/ericoliveiras/gate-guard/internal/model"
	"github.com/google/uuid"
)

type UserBuilder struct {
	ID         uuid.UUID
	Firstname  string
	Lastname   string
	DocumentId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (userBuilder *UserBuilder) SetID(id uuid.UUID) *UserBuilder {
	userBuilder.ID = id
	return userBuilder
}

func (userBuilder *UserBuilder) SetFirstname(firstname string) *UserBuilder {
	userBuilder.Firstname = firstname
	return userBuilder
}

func (userBuilder *UserBuilder) SetLastname(lastname string) *UserBuilder {
	userBuilder.Lastname = lastname
	return userBuilder
}

func (userBuilder *UserBuilder) SetDocumentId(documentId string) *UserBuilder {
	userBuilder.DocumentId = documentId
	return userBuilder
}

func (userBuilder *UserBuilder) SetCreatedAt(createdAt time.Time) *UserBuilder {
	userBuilder.CreatedAt = createdAt
	return userBuilder
}

func (userBuilder *UserBuilder) SetUpdatedAt(updatedAt time.Time) *UserBuilder {
	userBuilder.UpdatedAt = updatedAt
	return userBuilder
}

func (userBuilder *UserBuilder) Build() model.User {
	user := model.User{
		ID:         userBuilder.ID,
		FirstName:  userBuilder.Firstname,
		LastName:   userBuilder.Lastname,
		DocumentId: userBuilder.DocumentId,
		CreatedAt:  userBuilder.CreatedAt,
		UpdatedAt:  userBuilder.UpdatedAt,
	}

	return user
}
