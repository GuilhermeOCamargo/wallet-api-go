package services

type OwnerService interface {
}

type ownerServiceImpl struct {
}

func NewOwnerService() OwnerService {
	return &ownerServiceImpl{}
}
