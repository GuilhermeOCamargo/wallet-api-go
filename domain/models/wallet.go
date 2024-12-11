package models

type Wallet struct {
	Id                   string
	Owner                Owner
	BalanceAmountInCents int
}
