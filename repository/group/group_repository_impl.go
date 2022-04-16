package group

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type groupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepositoryImpl(db *gorm.DB) *groupRepositoryImpl {
	return &groupRepositoryImpl{db: db}
}

func (g *groupRepositoryImpl) Insert(ctx context.Context, group entity.Group) (err error) {
	if dbErr := g.db.WithContext(ctx).Create(&group).Error; dbErr != nil {
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

func (g *groupRepositoryImpl) InsertAll(ctx context.Context, groups []entity.Group) (err error) {
	if dbErr := g.db.WithContext(ctx).Create(&groups).Error; dbErr != nil {
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

func (g *groupRepositoryImpl) FindAll(ctx context.Context) (groups []entity.Group, err error) {
	if dbErr := g.db.WithContext(ctx).Preload("Address").Preload("Properties").Find(&groups).Error; dbErr != nil {
		log.Println(dbErr)
		err = repository.ErrDatabase
		return
	}
	return
}

func (g *groupRepositoryImpl) FindByID(ctx context.Context, id string) (group entity.Group, err error) {
	if dbErr := g.db.WithContext(ctx).Preload("Address").Preload("Properties").First(&group, "id = ?", id).Error; dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			err = repository.ErrRecordNotFound
			return
		}
		err = repository.ErrDatabase
	}
	return
}

func (g *groupRepositoryImpl) Update(ctx context.Context, group entity.Group) (err error) {
	if dbErr := g.db.WithContext(ctx).Save(&group).Error; dbErr != nil {
		err = repository.ErrRecordNotFound
	}
	return
}

func (g *groupRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	err = g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if result := tx.WithContext(ctx).Delete(&entity.Group{}, "id = ?", id); result.Error == nil {
			if result.RowsAffected < 1 {
				return repository.ErrRecordNotFound
			}
		} else {
			log.Println(result.Error)
			log.Println(result)
			return repository.ErrDatabase
		}

		if dbErr := tx.WithContext(ctx).Delete(&entity.Address{}, "id = ?", id).Error; dbErr != nil {
			log.Println(dbErr)
			return repository.ErrDatabase
		}

		if dbErr := tx.WithContext(ctx).Delete(&entity.Property{}, "group_id = ?", id).Error; dbErr != nil {
			log.Println(dbErr)
			return repository.ErrDatabase
		}

		return nil
	})

	return
}
