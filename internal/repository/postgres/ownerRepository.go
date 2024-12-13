package postgres

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"gorm.io/gorm"
)

type OwnerRepository interface {
	CreateOwner(e *entities.Owner) error
	FindOwnerByDocumentNumber(documentNumber string) (entities.Owner, error)
}

type ownerRepositoryImpl struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository {
	return &ownerRepositoryImpl{db: db}
}

func (w *ownerRepositoryImpl) CreateOwner(e *entities.Owner) error {
	log.Println("Repository", e)
	w.db.Create(&e)
	return nil
}

func (w *ownerRepositoryImpl) FindOwnerByDocumentNumber(documentNumber string) (entities.Owner, error) {
	var owner entities.Owner

	w.db.Where(&entities.Owner{DocumentNumber: documentNumber}).First(&owner)
	return owner, nil
}
