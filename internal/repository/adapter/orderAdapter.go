package adapter

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/postgres"
)

type OrderAdapter interface {
	CreateOrder(o *models.Order) error
}

type orderAdapterImpl struct {
	repository postgres.OrderRepository
}

func NewOrderAdapter(repository postgres.OrderRepository) OrderAdapter {
	return &orderAdapterImpl{
		repository: repository,
	}
}

func (a *orderAdapterImpl) CreateOrder(o *models.Order) error {
	entity := entities.NewOrder(o)

	a.repository.CreateOrder(entity)

	o.SetId(entity.ID)
	o.SetCreatedAt(entity.CreatedAt)
	o.SetUpdatedAt(entity.UpdatedAt)
	return nil
}
