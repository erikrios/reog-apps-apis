package showschedule

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
)

type ShowScheduleService interface {
	Create(ctx context.Context, p payload.CreateShowSchedule) (id string, err error)
	GetAll(ctx context.Context) (responses []response.ShowSchedule, err error)
	GetByID(ctx context.Context, id string) (response response.ShowScheduleDetails, err error)
	GetByGroupID(ctx context.Context, groupID string) (responses []response.ShowSchedule, err error)
	Update(ctx context.Context, id string, p payload.UpdateShowSchedule) (err error)
	Delete(ctx context.Context, id string) (err error)
}
