package payment

import (
	"context"
	"fmt"
)

type PaymentStore interface {
	CreatePayment(ctx context.Context)
}

type PaymentDBStore struct {
	db string
}

func NewPaymentStore(db string) PaymentStore {
	return &PaymentDBStore{db: db}
}

func (store *PaymentDBStore) CreatePayment(ctx context.Context) {
	fmt.Println("Creating Payment in the database", store.db)
}
