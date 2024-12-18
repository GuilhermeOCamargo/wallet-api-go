package service

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapter"
)

type TransactionService interface {
	CreateTransaction(t *models.Transaction) error
}

type transactionServiceImpl struct {
	adapter adapter.TransactionAdapter
}

func NewTransactionService(adapter adapter.TransactionAdapter) TransactionService {
	return &transactionServiceImpl{
		adapter: adapter,
	}
}

func (s *transactionServiceImpl) CreateTransaction(t *models.Transaction) error {
	err := s.adapter.CreateTransaction(t)
	if err != nil {
		return err
	}
	return nil
}
