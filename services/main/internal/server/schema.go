package server

type Hashed string

type UpdateUserBody struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
