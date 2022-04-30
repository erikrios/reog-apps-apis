package showschedule

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/showschedule"
)

type showScheduleServiceImpl struct {
	showScheduleRepository showschedule.ShowcheduleRepository
	groupRepository        group.GroupRepository
}

func NewShowScheduleServiceImpl(
	showScheduleRepository showschedule.ShowcheduleRepository,
	groupRepository group.GroupRepository,
) *showScheduleServiceImpl {
	return &showScheduleServiceImpl{
		showScheduleRepository: showScheduleRepository,
		groupRepository:        groupRepository,
	}
}

func (s *showScheduleServiceImpl) Create(ctx context.Context, p payload.CreateShowSchedule) (id string, err error) {
	return
}

func (s *showScheduleServiceImpl) GetAll(ctx context.Context) (responses []response.ShowSchedule, err error) {
	return
}

func (s *showScheduleServiceImpl) GetByID(ctx context.Context, id string) (response response.ShowScheduleDetails, err error) {
	return
}

func (s *showScheduleServiceImpl) GetByGroupID(ctx context.Context, groupID string) (responses []response.ShowSchedule, err error) {
	return
}

func (s *showScheduleServiceImpl) Update(ctx context.Context, id string, p payload.UpdateShowSchedule) (err error) {
	return
}

func (s *showScheduleServiceImpl) Delete(ctx context.Context, id string) (err error) { return }
