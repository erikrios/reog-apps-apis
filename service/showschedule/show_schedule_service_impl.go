package showschedule

import (
	"context"
	"time"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/showschedule"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"gopkg.in/validator.v2"
)

type showScheduleServiceImpl struct {
	showScheduleRepository showschedule.ShowScheduleRepository
	groupRepository        group.GroupRepository
	idGenerator            generator.IDGenerator
}

func NewShowScheduleServiceImpl(
	showScheduleRepository showschedule.ShowScheduleRepository,
	groupRepository group.GroupRepository,
	idGenerator generator.IDGenerator,
) *showScheduleServiceImpl {
	return &showScheduleServiceImpl{
		showScheduleRepository: showScheduleRepository,
		groupRepository:        groupRepository,
		idGenerator:            idGenerator,
	}
}

func (s *showScheduleServiceImpl) Create(ctx context.Context, p payload.CreateShowSchedule) (id string, err error) {
	if validateErr := validator.Validate(p); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	startOn, parseErr := time.Parse(time.RFC822, p.StartOn)
	if parseErr != nil {
		err = service.ErrTimeParsing
		return
	}

	finishOn, parseErr := time.Parse(time.RFC822, p.FinishOn)
	if parseErr != nil {
		err = service.ErrTimeParsing
		return
	}

	if _, repoErr := s.groupRepository.FindByID(ctx, p.GroupID); repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	id, genErr := s.idGenerator.GenerateShowScheduleID()
	if genErr != nil {
		err = service.MapError(genErr)
		return
	}

	showSchedule := entity.ShowSchedule{
		ID:       id,
		GroupID:  p.GroupID,
		Place:    p.Place,
		StartOn:  startOn,
		FinishOn: finishOn,
	}

	if repoErr := s.showScheduleRepository.Insert(ctx, showSchedule); repoErr != nil {
		err = service.MapError(repoErr)
	}

	return
}

func (s *showScheduleServiceImpl) GetAll(ctx context.Context) (responses []response.ShowSchedule, err error) {
	entities, repoErr := s.showScheduleRepository.FindAll(ctx)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	responses = make([]response.ShowSchedule, 0)

	for _, entity := range entities {
		response := response.ShowSchedule{
			ID:       entity.ID,
			GroupID:  entity.GroupID,
			Place:    entity.Place,
			StartOn:  entity.StartOn.Format(time.RFC822),
			FinishOn: entity.FinishOn.Format(time.RFC822),
		}

		responses = append(responses, response)
	}

	return
}

func (s *showScheduleServiceImpl) GetByID(ctx context.Context, id string) (response response.ShowScheduleDetails, err error) {
	entity, repoErr := s.showScheduleRepository.FindByID(ctx, id)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	response.ID = entity.ID
	response.Place = entity.Place
	response.StartOn = entity.StartOn.Format(time.RFC822)
	response.FinishOn = entity.FinishOn.Format(time.RFC822)

	groupEntity, repoErr := s.groupRepository.FindByID(ctx, entity.GroupID)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	response.GroupID = groupEntity.ID
	response.GroupName = groupEntity.Name

	return
}

func (s *showScheduleServiceImpl) GetByGroupID(ctx context.Context, groupID string) (responses []response.ShowSchedule, err error) {
	_, repoErr := s.groupRepository.FindByID(ctx, groupID)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	entities, repoErr := s.showScheduleRepository.FindByGroupID(ctx, groupID)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	responses = make([]response.ShowSchedule, 0)

	for _, entity := range entities {
		response := response.ShowSchedule{
			ID:       entity.ID,
			GroupID:  entity.GroupID,
			Place:    entity.Place,
			StartOn:  entity.StartOn.Format(time.RFC822),
			FinishOn: entity.FinishOn.Format(time.RFC822),
		}

		responses = append(responses, response)
	}

	return
}

func (s *showScheduleServiceImpl) Update(ctx context.Context, id string, p payload.UpdateShowSchedule) (err error) {
	return
}

func (s *showScheduleServiceImpl) Delete(ctx context.Context, id string) (err error) { return }
