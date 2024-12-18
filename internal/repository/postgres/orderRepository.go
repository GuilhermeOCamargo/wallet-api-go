package postgres

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(e *entities.OrderEntity) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) CreateOrder(e *entities.OrderEntity) error {
	r.db.Create(&e)
	return nil
}
