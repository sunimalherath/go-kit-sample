package account

import "context"

type User struct {
	ID       string `json:"id, omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository - methods in this type is to interface with the database
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}