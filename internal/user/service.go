package user

import (
	"context"
	"fmt"
)

type User struct {
	Email string `json:"email"`
}

type Service interface {
	GetUser(ctx context.Context, email string) (User, error)
}

type UserService struct {
	store UserStore
}

func NewService(store UserStore) Service {
	return &UserService{store: store}
}

func (s *UserService) GetUser(ctx context.Context, email string) (User, error) {
	fmt.Println("Finding user in the service...")

	s.store.FindUserById(context.TODO(), "123")
	return User{Email: email}, nil
}
