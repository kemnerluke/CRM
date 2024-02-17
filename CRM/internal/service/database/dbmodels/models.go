package database

type Customer struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber int    `json:"phone_number"`
	Email       string `json:"email"`
}