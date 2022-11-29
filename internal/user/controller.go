package user

import (
	"context"
	"fmt"
	"net/http"
)

type Controller interface {
	GetUser() func(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service Service
}

// Can accept any service, UserService, PaymentService
func NewController(s Service) Controller {
	return &controller{service: s}
}

func (c *controller) GetUser() func(w http.ResponseWriter, r *http.Request) {
	// c.service.GetUser()
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Finding user in the controller...")
		c.service.GetUser(context.Background(), "perica")

		w.WriteHeader(200)
		w.Write([]byte("Status: OK\n"))
	}
}
