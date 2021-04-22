package account

import "context"

//Service - interface to define methods to expose to transport & use the interface to implement business logic
// these methods will be expose from the microservice.
type Service interface {
	 CreateUser(ctx context.Context, email string, password string) (string, error)
	 GetUser(ctx context.Context, id string) (string, error)
}