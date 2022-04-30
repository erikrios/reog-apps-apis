package showschedule

import (
	"context"
	"errors"
	"log"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type showScheduleRepositoryImpl struct {
	db *gorm.DB
}

func NewShowScheduleRepositoryImpl(db *gorm.DB) *showScheduleRepositoryImpl {
	return &showScheduleRepositoryImpl{db: db}
}

func (s *showScheduleRepositoryImpl) Insert(ctx context.Context, showSchedule entity.ShowSchedule) (err error) {
	if dbErr := s.db.WithContext(ctx).Create(&showSchedule).Error; dbErr != nil {
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

func (s *showScheduleRepositoryImpl) FindAll(ctx context.Context) (showSchedules []entity.ShowSchedule, err error) {
	if dbErr := s.db.WithContext(ctx).Find(&showSchedules).Error; dbErr != nil {
		log.Println(dbErr)
		err = repository.ErrDatabase
		return
	}
	return
}

func (s *showScheduleRepositoryImpl) FindByGroupID(ctx context.Context, groupID string) (showSchedules []entity.ShowSchedule, err error) {
	if dbErr := s.db.WithContext(ctx).Where("group_id = ?", groupID).Find(&showSchedules).Error; dbErr != nil {
		log.Println(dbErr)
		err = repository.ErrDatabase
		return
	}
	return
}

func (s *showScheduleRepositoryImpl) Update(ctx context.Context, id string, showSchedule entity.ShowSchedule) (err error) {
	if result := s.db.WithContext(ctx).Where("id = ?", id).UpdateColumns(&showSchedule); result.Error != nil {
		log.Println(result.Error)
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}

func (s *showScheduleRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	if result := s.db.WithContext(ctx).Delete(&entity.ShowSchedule{}, "id = ?", id); result.Error != nil {
		log.Println(result.Error)
		err = repository.ErrDatabase
	} else {
		if result.RowsAffected < 1 {
			err = repository.ErrRecordNotFound
		}
	}
	return
}
