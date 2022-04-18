package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/address"
	"github.com/erikrios/reog-apps-apis/repository/village"
	"github.com/erikrios/reog-apps-apis/service"
	"gopkg.in/validator.v2"
)

type addressServiceImpl struct {
	addressRepository address.AddressRepository
	villageRepository village.VillageRepository
}

func NewAddressServiceImpl(
	addressRepository address.AddressRepository,
	villageRepository village.VillageRepository,
) *addressServiceImpl {
	return &addressServiceImpl{
		addressRepository: addressRepository,
		villageRepository: villageRepository,
	}
}

func (a *addressServiceImpl) Update(ctx context.Context, id string, p payload.UpdateAddress) (err error) {
	if validateErr := validator.Validate(p); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	village, villageErr := a.villageRepository.FindByID(p.VillageID)
	if villageErr != nil {
		err = service.MapError(villageErr)
		return
	}

	address := entity.Address{
		ID:           id,
		Address:      p.Address,
		VillageID:    village.ID,
		VillageName:  village.Name,
		DistrictID:   village.District.ID,
		DistrictName: village.District.Name,
		RegencyID:    village.District.Regency.ID,
		RegencyName:  village.District.Regency.Name,
		ProvinceID:   village.District.Regency.Province.ID,
		ProvinceName: village.District.Regency.Province.Name,
	}

	repoErr := a.addressRepository.Update(ctx, id, address)
	if repoErr != nil {
		err = service.MapError(repoErr)
	}
	return
}
