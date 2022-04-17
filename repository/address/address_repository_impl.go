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

func (a *addressRepositoryImpl) Update(ctx context.Context, id string, address entity.Address) (err error) {
	if result := a.db.WithContext(ctx).Where("id = ?", id).UpdateColumns(&address); result.Error != nil {
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}
