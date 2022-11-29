package user

import (
	"context"
	"fmt"
)

type UserStore interface {
	FindUserById(ctx context.Context, userId string) (User, error)
}

type UserDBStore struct {
	db string
}

func NewUserStore(db string) UserStore {
	return &UserDBStore{db: db}
}

func (store UserDBStore) FindUserById(ctx context.Context, userId string) (User, error) {
	fmt.Println("Finding user in the Database...", store.db)
	return User{}, nil
}
