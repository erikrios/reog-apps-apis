package group

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/village"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"gopkg.in/validator.v2"
)

type groupServiceImpl struct {
	groupRepository   group.GroupRepository
	villageRepository village.VillageRepository
}

func NewGroupServiceImpl(
	groupRepository group.GroupRepository,
	villageRepository village.VillageRepository,
) *groupServiceImpl {
	return &groupServiceImpl{
		groupRepository:   groupRepository,
		villageRepository: villageRepository,
	}
}

func (g *groupServiceImpl) Create(ctx context.Context, p payload.CreateGroup) (id string, err error) {
	if validateErr := validator.Validate(p); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	village, villageErr := g.villageRepository.FindByID(p.VillageID)
	if villageErr != nil {
		err = service.MapError(villageErr)
		return
	}

	id, genErr := generator.GenerateGroupID()
	if genErr != nil {
		err = service.MapError(genErr)
		return
	}

	group := entity.Group{
		ID:     id,
		Name:   p.Name,
		Leader: p.Leader,
		Address: entity.Address{
			ID:           id,
			Address:      p.Address,
			VillageID:    village.ID,
			VillageName:  village.Name,
			DistrictID:   village.District.ID,
			DistrictName: village.District.Name,
			RegencyID:    village.District.Regency.ID,
			RegencyName:  village.District.Regency.Name,
			ProvinceID:   village.District.Regency.Province.ID,
			ProvinceName: village.District.Regency.Province.Name,
		},
	}

	if repoErr := g.groupRepository.Insert(ctx, group); repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	return
}

func (g *groupServiceImpl) GetAll(ctx context.Context) (responses []response.Group, err error) {
	return
}

func (g *groupServiceImpl) GetByID(ctx context.Context, id string) (response response.Group, err error) {
	return
}

func (g *groupServiceImpl) Update(ctx context.Context, id string, p payload.UpdateGroup) (err error) {
	return
}

func (g *groupServiceImpl) Delete(ctx context.Context, id string) (err error) { return }

func (g *groupServiceImpl) GeterateQRCode(ctx context.Context, id string) (file []byte, err error) {
	return
}
