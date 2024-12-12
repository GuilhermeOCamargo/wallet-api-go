package adapters

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/entities"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/postgres"
)

type OwnerAdapter interface {
	CreateOwner(o *models.Owner)
}

type ownerAdapterImpl struct {
	repository postgres.OwnerRepository
}

func NewOwnerAdapter(repository postgres.OwnerRepository) OwnerAdapter {
	return &ownerAdapterImpl{
		repository: repository,
	}
}

func (a ownerAdapterImpl) CreateOwner(o *models.Owner) {
	entity := entities.NewOwner(o)

	a.repository.CreateOwner(entity)
	log.Println("Owner criado no banco de dados", entity)

	o.SetId(entity.ID)
	o.SetCreatedAt(entity.CreatedAt)
	o.SetUpdatedAt(entity.UpdatedAt)
}
