package admin

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepositoryImpl(db *gorm.DB) *adminRepositoryImpl {
	return &adminRepositoryImpl{db: db}
}

func (a *adminRepositoryImpl) FindByUsername(ctx context.Context, username string) (admin entity.Admin, err error) {
	if dbErr := a.db.First(&admin, "username = ?", username).Error; dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			err = repository.ErrRecordNotFound
			return
		}
		err = repository.ErrDatabase
		log.Println(err)
	}
	return
}
