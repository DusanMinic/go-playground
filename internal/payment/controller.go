package payment

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentService interface {
	CreatePayment(ctx context.Context) (*Payment, error)
}

type User struct {
	Email string `json:"email"`
}

type UserService interface {
	GetUser(ctx context.Context, email string) (*User, error)
}

type controller struct {
	paymentService PaymentService
	userService    UserService
}

func NewController(paymentService PaymentService, userService UserService) *controller {
	return &controller{
		paymentService: paymentService,
		userService:    userService,
	}
}

func (c *controller) CreatePayment() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create Payment controller")

		c.userService.GetUser(context.Background(), "perica")

		c.paymentService.CreatePayment(context.Background())

		w.WriteHeader(200)
		w.Write([]byte("Status: OK\n"))
	}
}
