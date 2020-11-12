package model

// User db model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"first_name"`
	LName string `json:"last_name"`
	Email string `json:"email"`
}
