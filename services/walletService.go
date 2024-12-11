package services

type WalletService interface {
}

type walletServiceImpl struct {
}

func NewWalletService() WalletService {
	return &walletServiceImpl{}
}
