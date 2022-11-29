package main

import (
	PaymentService "api/internal/payment"
	UserService "api/internal/user"
	"log"
	"net/http"
)

func main() {
	db := "postgres"

	userService := UserService.NewService(UserService.NewUserStore(db))
	paymentService := PaymentService.NewService(PaymentService.NewPaymentStore(db))

	userController := UserService.NewController(userService)
	paymentController := PaymentService.NewController(paymentService, userService)

	http.HandleFunc("/user", userController.GetUser())
	http.HandleFunc("/payment", paymentController.CreatePayment())
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
