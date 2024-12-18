package models

import (
	"time"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/enums"
)

type Transaction struct {
	id            uint
	wallet        *Wallet
	amountInCents int
	operation     *enums.Operation
	order         *Order
	createdAt     time.Time
	updatedAt     time.Time
}

func NewTransaction(wallet *Wallet, amountInCents int, operation *enums.Operation, order *Order) *Transaction {
	return &Transaction{
		wallet:        wallet,
		amountInCents: amountInCents,
		operation:     operation,
		order:         order,
	}
}

func (t *Transaction) Id() uint {
	return t.id
}

func (t *Transaction) Wallet() *Wallet {
	return t.wallet
}

func (t *Transaction) AmountInCents() int {
	return t.amountInCents
}

func (t *Transaction) Operation() *enums.Operation {
	return t.operation
}

func (t *Transaction) Order() *Order {
	return t.order
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Transaction) SetId(id uint) {
	t.id = id
}

func (t *Transaction) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Transaction) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
