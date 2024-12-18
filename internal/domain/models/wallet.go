package models

import "time"

type Wallet struct {
	id                   uint
	owner                *Owner
	balanceAmountInCents int
	createdAt            time.Time
	updatedAt            time.Time
}

func NewWalletWithOwnerAndBalance(owner *Owner, balanceAmountInCents int) *Wallet {
	return &Wallet{
		owner:                owner,
		balanceAmountInCents: balanceAmountInCents,
	}
}

func NewWallet(owner *Owner, balanceAmountInCents int, id uint, createdAt, updatedAt time.Time) *Wallet {
	return &Wallet{
		owner:                owner,
		balanceAmountInCents: balanceAmountInCents,
		id:                   id,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
	}
}

func (w Wallet) Id() uint {
	return w.id
}

func (w Wallet) Owner() *Owner {
	return w.owner
}

func (w Wallet) BalanceAmountInCents() int {
	return w.balanceAmountInCents
}

func (w *Wallet) SetId(id uint) {
	w.id = id
}

func (w *Wallet) SetCreatedAt(createdAt time.Time) {
	w.createdAt = createdAt
}

func (w *Wallet) SetUpdatedAt(updatedAt time.Time) {
	w.updatedAt = updatedAt
}
