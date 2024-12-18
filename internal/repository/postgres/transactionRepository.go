package postgres

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateOrder(e *entities.TransactionEntity) error
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) CreateOrder(e *entities.TransactionEntity) error {
	r.db.Create(&e)
	return nil
}
