package adapter

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/postgres"
)

type TransactionAdapter interface {
	CreateTransaction(t *models.Transaction) error
}

type transactionAdapterImpl struct {
	repository postgres.TransactionRepository
}

func NewTransactionAdapter(repository postgres.TransactionRepository) TransactionAdapter {
	return &transactionAdapterImpl{
		repository: repository,
	}
}

func (a *transactionAdapterImpl) CreateTransaction(t *models.Transaction) error {
	entity := entities.NewTransactionEntity(t)

	a.repository.CreateOrder(entity)

	t.SetId(entity.ID)
	t.SetCreatedAt(entity.CreatedAt)
	t.SetUpdatedAt(entity.UpdatedAt)
	return nil
}
