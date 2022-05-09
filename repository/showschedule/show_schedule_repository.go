package showschedule

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type ShowScheduleRepository interface {
	Insert(ctx context.Context, showSchedule entity.ShowSchedule) (err error)
	FindAll(ctx context.Context) (showSchedules []entity.ShowSchedule, err error)
	FindByID(ctx context.Context, id string) (showSchedule entity.ShowSchedule, err error)
	FindByGroupID(ctx context.Context, groupID string) (showSchedules []entity.ShowSchedule, err error)
	Update(ctx context.Context, id string, showSchedule entity.ShowSchedule) (err error)
	Delete(ctx context.Context, id string) (err error)
}
