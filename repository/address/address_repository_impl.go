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
	if dbErr := a.db.WithContext(ctx).Save(&address).Error; dbErr != nil {
		err = repository.ErrRecordNotFound
	}
	return
}
