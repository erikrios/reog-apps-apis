package group

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/erikrios/reog-apps-apis/utils/logging"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type groupRepositoryImpl struct {
	db     *gorm.DB
	logger logging.Logging
}

func NewGroupRepositoryImpl(db *gorm.DB, logger logging.Logging) *groupRepositoryImpl {
	return &groupRepositoryImpl{db: db, logger: logger}
}

func (g *groupRepositoryImpl) Insert(ctx context.Context, group entity.Group) (err error) {
	if dbErr := g.db.WithContext(ctx).Create(&group).Error; dbErr != nil {
		var pqErr *pgconn.PgError
		if ok := errors.As(dbErr, &pqErr); ok && pqErr.Code == "23505" {
			err = repository.ErrRecordAlreadyExists
			return
		}

		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(g.logger, dbErr.Error())

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

		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(g.logger, dbErr.Error())

		log.Println(dbErr)
		err = repository.ErrDatabase
	}
	return
}

func (g *groupRepositoryImpl) FindAll(ctx context.Context) (groups []entity.Group, err error) {
	if dbErr := g.db.WithContext(ctx).Preload("Address").Preload("Properties").Find(&groups).Error; dbErr != nil {
		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(g.logger, dbErr.Error())

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

		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(g.logger, dbErr.Error())

		err = repository.ErrDatabase
	}
	return
}

func (g *groupRepositoryImpl) Update(ctx context.Context, id string, group entity.Group) (err error) {
	if result := g.db.WithContext(ctx).Where("id = ?", id).UpdateColumns(&group); result.Error != nil {
		go func(logger logging.Logging, message string) {
			logger.Error(message)
		}(g.logger, result.Error.Error())

		log.Println(result.Error)
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
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
			go func(logger logging.Logging, message string) {
				logger.Error(message)
			}(g.logger, result.Error.Error())

			log.Println(result.Error)
			return repository.ErrDatabase
		}

		if dbErr := tx.WithContext(ctx).Delete(&entity.Address{}, "id = ?", id).Error; dbErr != nil {
			go func(logger logging.Logging, message string) {
				logger.Error(message)
			}(g.logger, dbErr.Error())

			log.Println(dbErr)
			return repository.ErrDatabase
		}

		if dbErr := tx.WithContext(ctx).Delete(&entity.Property{}, "group_id = ?", id).Error; dbErr != nil {
			go func(logger logging.Logging, message string) {
				logger.Error(message)
			}(g.logger, dbErr.Error())

			log.Println(dbErr)
			return repository.ErrDatabase
		}

		return nil
	})

	return
}
