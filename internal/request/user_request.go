package request

type CreateUser struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	DocumentId string `json:"document_id"`
}
