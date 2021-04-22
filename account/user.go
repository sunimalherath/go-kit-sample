package account

import "context"

// User - to represent a user inside of the business logic
type User struct {
	ID       string `json:"id, omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository - methods in this type is to interface with the database
// The functions are very similar to the ones that are in Service interface.
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}