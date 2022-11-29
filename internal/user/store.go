package user

import (
	"context"
	"fmt"
)

type UserDBStore struct {
	db string
}

func NewUserStore(db string) *UserDBStore {
	return &UserDBStore{db: db}
}

func (store UserDBStore) FindUserById(ctx context.Context, userId string) (User, error) {
	fmt.Println("Finding user in the Database...", store.db)
	return User{}, nil
}
