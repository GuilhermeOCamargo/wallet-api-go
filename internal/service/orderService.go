package service

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapter"
)

type OrderService interface {
	CreateOrder(o *models.Order) error
}

type orderServiceImpl struct {
	adapter adapter.OrderAdapter
}

func NewOrderService(adapter adapter.OrderAdapter) OrderService {
	return &orderServiceImpl{
		adapter: adapter,
	}
}

func (s *orderServiceImpl) CreateOrder(o *models.Order) error {
	err := s.adapter.CreateOrder(o)
	if err != nil {
		return err
	}
	return nil
}
