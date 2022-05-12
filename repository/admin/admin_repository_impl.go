package admin

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/erikrios/reog-apps-apis/utils/logging"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db     *gorm.DB
	logger logging.Logging
}

func NewAdminRepositoryImpl(db *gorm.DB, logger logging.Logging) *adminRepositoryImpl {
	return &adminRepositoryImpl{db: db, logger: logger}
}

func (a *adminRepositoryImpl) FindByUsername(ctx context.Context, username string) (admin entity.Admin, err error) {
	if dbErr := a.db.WithContext(ctx).First(&admin, "username = ?", username).Error; dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			err = repository.ErrRecordNotFound
			return
		}

		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(a.logger, dbErr.Error())

		err = repository.ErrDatabase
		log.Println(dbErr)
	}
	return
}
