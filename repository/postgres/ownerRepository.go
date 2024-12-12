package postgres

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/repository/entities"
	"gorm.io/gorm"
)

type OwnerRepository interface {
	CreateOwner(e *entities.Owner)
}

type ownerRepositoryImpl struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository {
	return &ownerRepositoryImpl{db: db}
}

func (w *ownerRepositoryImpl) CreateOwner(e *entities.Owner) {
	log.Println("Repository", e)
	w.db.Create(&e)
}
