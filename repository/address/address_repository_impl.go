package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"gorm.io/gorm"
)

type addressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepositoryImpl(db *gorm.DB) *addressRepositoryImpl {
	return &addressRepositoryImpl{db: db}
}

func (a *addressRepositoryImpl) Update(ctx context.Context, address entity.Address) (err error) {
	if result := a.db.WithContext(ctx).Model(&address).Updates(entity.Address{
		Address:      address.Address,
		VillageID:    address.VillageID,
		VillageName:  address.VillageName,
		DistrictID:   address.DistrictID,
		DistrictName: address.DistrictName,
		RegencyID:    address.RegencyID,
		RegencyName:  address.RegencyName,
		ProvinceID:   address.ProvinceID,
		ProvinceName: address.ProvinceName,
	}); result.Error != nil {
		err = repository.ErrRecordNotFound
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}
