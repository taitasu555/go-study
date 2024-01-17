package usecase

import (
	"context"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)



type addressUsecase struct {
	addressRepository domain.AddressRepository
	contextTimeout time.Duration
}


func NewAddressUsecase(addressRepository domain.AddressRepository, timeout time.Duration) domain.AddressUsecase {
	return &addressUsecase{
		addressRepository: addressRepository,
		contextTimeout: timeout,
	}
}


func (tu *addressUsecase) Create(c context.Context, address *domain.Address) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.Create(ctx, address)
}


func (tu *addressUsecase) FetchAddressByUserID (c context.Context, userID string) ([]domain.Address, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.FetchAddressByUserID(ctx, userID)
}
