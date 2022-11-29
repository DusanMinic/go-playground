package payment

import (
	"context"
	"fmt"
	"net/http"
)

type Controller interface {
	CreatePayment() func(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	paymentService PaymentService
	userService    UserService
}

func NewController(paymentService PaymentService, userService UserService) Controller {
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
