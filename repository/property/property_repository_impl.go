package property

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type propertyRepositoryImpl struct {
	db *gorm.DB
}

func NewPropertyRepositoryImpl(db *gorm.DB) *propertyRepositoryImpl {
	return &propertyRepositoryImpl{db: db}
}

func (p *propertyRepositoryImpl) Insert(ctx context.Context, property entity.Property) (err error) {
	if dbErr := p.db.WithContext(ctx).Create(&property).Error; dbErr != nil {
		var pqErr *pgconn.PgError
		if ok := errors.As(dbErr, &pqErr); ok && pqErr.Code == "23505" {
			err = repository.ErrRecordAlreadyExists
			return
		}
		log.Println(dbErr)
		err = repository.ErrDatabase
	}
	return
}

func (p *propertyRepositoryImpl) Update(ctx context.Context, property entity.Property) (err error) {
	if result := p.db.WithContext(ctx).Model(&property).Updates(entity.Property{
		Name:        property.Name,
		Description: property.Description,
		Amount:      property.Amount,
	}); result.Error != nil {
		log.Println(result.Error)
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}

func (p *propertyRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	if result := p.db.WithContext(ctx).Delete(&entity.Property{}, "id = ?", id); result.Error != nil {
		log.Println(result.Error)
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}
