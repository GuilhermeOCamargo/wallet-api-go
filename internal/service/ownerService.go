package service

import (
	"errors"

	apperrors "github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapter"
)

type OwnerService interface {
	CreateOwner(o *models.Owner) error
}

type ownerServiceImpl struct {
	adapter adapter.OwnerAdapter
}

func NewOwnerService(adapter adapter.OwnerAdapter) OwnerService {
	return &ownerServiceImpl{
		adapter: adapter,
	}
}

func (s *ownerServiceImpl) CreateOwner(o *models.Owner) error {
	if s.canCreateOwner(o) {
		s.adapter.CreateOwner(o)
		return nil
	}
	return errors.New("ja existe um owner cadastrado com as informações")
}

func (s *ownerServiceImpl) findOwnerByDocumentNumber(documentNumber string) (*models.Owner, error) {
	owner, err := s.adapter.FindOwnerByDocumentNumber(documentNumber)
	if err != nil {
		return &models.Owner{}, apperrors.NewCodeError(500, "Erro ao buscar owner")
	}
	if owner.Id() == 0 {
		return &models.Owner{}, apperrors.NewCodeError(404, "Owner nao encontrado")
	}
	return owner, nil
}

func (s *ownerServiceImpl) canCreateOwner(o *models.Owner) bool {
	_, err := s.findOwnerByDocumentNumber(o.DocumentNumber())
	var ce *apperrors.CodeError
	if err != nil && errors.As(err, &ce) {
		return true
	}
	return false
}
