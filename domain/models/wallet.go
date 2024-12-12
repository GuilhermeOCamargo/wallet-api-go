package models

type Wallet struct {
	id                   string
	owner                Owner
	balanceAmountInCents int
}

func NewWallet(owner Owner, balanceAmountInCents int) Wallet {
	return Wallet{
		owner:                owner,
		balanceAmountInCents: balanceAmountInCents,
	}
}

func (w Wallet) GetId() string {
	return w.id
}

func (w Wallet) GetOwner() Owner {
	return w.owner
}

func (w Wallet) GetBalanceAmountInCents() int {
	return w.balanceAmountInCents
}
