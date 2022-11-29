package payment

import (
	"context"
	"fmt"
)

type PaymentDBStore struct {
	db string
}

func NewPaymentStore(db string) *PaymentDBStore {
	return &PaymentDBStore{db: db}
}

func (store *PaymentDBStore) CreatePayment(ctx context.Context) {
	fmt.Println("Creating Payment in the database", store.db)
}
