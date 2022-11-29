package payment

import (
	"context"
	"fmt"
)

// Create Service interface in service.go
// Create Payment struct in service.go
// Create newService func which returns Service interface
// implement receiver func PaymentService (from Interface)

type Payment struct {
	Amount string `json:"amount"`
}

type PaymentService interface {
	CreatePayment(ctx context.Context) (Payment, error)
}

type User struct {
	Email string `json:"email"`
}

type UserService interface {
	GetUser(ctx context.Context, email string) (User, error)
}

type paymentService struct {
	store PaymentStore
}

func NewService(store PaymentStore) PaymentService {
	return &paymentService{store: store}
}

func (s *paymentService) CreatePayment(ctx context.Context) (Payment, error) {
	fmt.Println("Creating Payment service")

	s.store.CreatePayment(ctx)

	return Payment{}, nil
}
