package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/erikrios/reog-apps-apis/utils/logging"
	"gorm.io/gorm"
)

type addressRepositoryImpl struct {
	db     *gorm.DB
	logger logging.Logging
}

func NewAddressRepositoryImpl(db *gorm.DB, logger logging.Logging) *addressRepositoryImpl {
	return &addressRepositoryImpl{db: db, logger: logger}
}

func (a *addressRepositoryImpl) Update(ctx context.Context, id string, address entity.Address) (err error) {
	if result := a.db.WithContext(ctx).Where("id = ?", id).UpdateColumns(&address); result.Error != nil {
		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(a.logger, result.Error.Error())

		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}
