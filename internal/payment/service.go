package payment

import (
	"context"
	"fmt"
)

type Payment struct {
	Amount string `json:"amount"`
}

type PaymentStore interface {
	CreatePayment(ctx context.Context)
}

type paymentService struct {
	store PaymentStore
}

func NewService(store PaymentStore) *paymentService {
	return &paymentService{store: store}
}

func (s *paymentService) CreatePayment(ctx context.Context) (*Payment, error) {
	fmt.Println("Creating Payment service")

	s.store.CreatePayment(ctx)

	return &Payment{}, nil
}
