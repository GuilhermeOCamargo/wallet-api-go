package models

import (
	"time"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/enums"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/vo"
)

type Order struct {
	id                        uint
	debitWallet, creditWallet *Wallet
	createdAt, updatedAt      time.Time
	title, description        string
	amountInCents             int
	transactions              []*Transaction
	orderType                 *enums.OrderType
}

func NewOrder(vo *vo.CreateOrderVo, debitWallet, creditWallet *Wallet, orderType *enums.OrderType) *Order {
	return &Order{
		title:         vo.Title(),
		description:   vo.Description(),
		amountInCents: vo.AmountInCents(),
		debitWallet:   debitWallet,
		creditWallet:  creditWallet,
		orderType:     orderType,
	}
}

func (o *Order) Id() uint {
	return o.id
}

func (o *Order) DebitWallet() *Wallet {
	return o.debitWallet
}

func (o *Order) CreditWallet() *Wallet {
	return o.creditWallet
}

func (o *Order) Title() string {
	return o.title
}

func (o *Order) Description() string {
	return o.description
}

func (o *Order) AmountInCents() int {
	return o.amountInCents
}

func (o *Order) Transactions() []*Transaction {
	return o.transactions
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

func (o *Order) OrderType() *enums.OrderType {
	return o.orderType
}

func (o *Order) SetId(id uint) {
	o.id = id
}

func (o *Order) SetCreatedAt(createdAt time.Time) {
	o.createdAt = createdAt
}

func (o *Order) SetUpdatedAt(updatedAt time.Time) {
	o.updatedAt = updatedAt
}

func (o *Order) AddTransaction(t ...*Transaction) {
	o.transactions = append(o.transactions, t...)
}
