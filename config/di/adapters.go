package di

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapter"

type adapters struct {
	ownerAdapter       adapter.OwnerAdapter
	walletAdapter      adapter.WalletAdapter
	orderAdapter       adapter.OrderAdapter
	transactionAdapter adapter.TransactionAdapter
}

func InitializeAdapters(repos repositories) adapters {
	return adapters{
		ownerAdapter:       adapter.NewOwnerAdapter(repos.ownerRepo),
		walletAdapter:      adapter.NewWalletAdapter(repos.walletRepo),
		orderAdapter:       adapter.NewOrderAdapter(repos.orderRepo),
		transactionAdapter: adapter.NewTransactionAdapter(repos.transactionRepo),
	}
}
