package di

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/config/database"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/postgres"
)

type repositories struct {
	ownerRepo       postgres.OwnerRepository
	walletRepo      postgres.WalletRepository
	orderRepo       postgres.OrderRepository
	transactionRepo postgres.TransactionRepository
}

func InitializeRepositories() repositories {
	db := database.DB
	return repositories{
		ownerRepo:       postgres.NewOwnerRepository(db),
		walletRepo:      postgres.NewWalletRepository(db),
		orderRepo:       postgres.NewOrderRepository(db),
		transactionRepo: postgres.NewTransactionRepository(db),
	}
}
