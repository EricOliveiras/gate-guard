package model

import (
	"time"
)

type User struct {
	ID         string    `db:"id,uuid,pk"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	DocumentId string    `db:"document_id,unique"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
