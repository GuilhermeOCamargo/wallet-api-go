package entities

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Name           string
	DocumentNumber string
}

func NewOwner(o *models.Owner) *Owner {
	return &Owner{
		Name:           o.Name(),
		DocumentNumber: o.DocumentNumber(),
	}
}
