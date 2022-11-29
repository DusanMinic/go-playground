package user

import (
	"context"
	"fmt"
)

type UserStore interface {
	FindUserById(ctx context.Context, userId string) (User, error)
}

type User struct {
	Email string `json:"email"`
}

type UserService struct {
	store UserStore
}

func NewService(store UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUser(ctx context.Context, email string) (*User, error) {
	fmt.Println("Finding user in the service...")

	s.store.FindUserById(context.TODO(), "123")
	return &User{Email: email}, nil
}
