package user

import (
	"context"
	"fmt"
	"net/http"
)

type Service interface {
	GetUser(ctx context.Context, email string) (*User, error)
}

type controller struct {
	service Service
}

func NewController(s Service) *controller {
	return &controller{service: s}
}

func (c *controller) GetUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Finding user in the controller...")
		c.service.GetUser(context.Background(), "perica")

		w.WriteHeader(200)
		w.Write([]byte("Status: OK\n"))
	}
}
