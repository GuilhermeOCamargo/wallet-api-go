package services

import (
	"errors"

	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/adapters"
)

type OwnerService interface {
	CreateOwner(o *models.Owner) error
}

type ownerServiceImpl struct {
	adapter adapters.OwnerAdapter
}

func NewOwnerService(adapter adapters.OwnerAdapter) OwnerService {
	return &ownerServiceImpl{
		adapter: adapter,
	}
}

func (s *ownerServiceImpl) CreateOwner(o *models.Owner) error {
	if canCreateOwner(o) {
		s.adapter.CreateOwner(o)
		return nil
	}
	return errors.New("ja existe um owner cadastrado com as informações")
}

func canCreateOwner(o *models.Owner) bool {
	return true
}
