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
	"github.com/skip2/go-qrcode"
	"gopkg.in/validator.v2"
)

type groupServiceImpl struct {
	groupRepository   group.GroupRepository
	villageRepository village.VillageRepository
	idGenerator       generator.IDGenerator
	qrCodeGenerator   generator.QRCodeGenerator
}

func NewGroupServiceImpl(
	groupRepository group.GroupRepository,
	villageRepository village.VillageRepository,
	idGenerator generator.IDGenerator,
	qrCodeGenerator generator.QRCodeGenerator,
) *groupServiceImpl {
	return &groupServiceImpl{
		groupRepository:   groupRepository,
		villageRepository: villageRepository,
		idGenerator:       idGenerator,
		qrCodeGenerator:   qrCodeGenerator,
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

	id, genErr := g.idGenerator.GenerateGroupID()
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
	groups, repoErr := g.groupRepository.FindAll(ctx)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	responses = mapToModels(groups)
	return
}

func (g *groupServiceImpl) GetByID(ctx context.Context, id string) (response response.Group, err error) {
	group, repoErr := g.groupRepository.FindByID(ctx, id)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	response = mapToModel(group)
	return
}

func (g *groupServiceImpl) Update(ctx context.Context, id string, p payload.UpdateGroup) (err error) {
	if validateErr := validator.Validate(p); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	group := entity.Group{
		ID:     id,
		Name:   p.Name,
		Leader: p.Leader,
	}

	if repoErr := g.groupRepository.Update(ctx, id, group); repoErr != nil {
		err = service.MapError(repoErr)
	}
	return
}

func (g *groupServiceImpl) Delete(ctx context.Context, id string) (err error) {
	if repoErr := g.groupRepository.Delete(ctx, id); repoErr != nil {
		err = service.MapError(repoErr)
	}
	return
}

func (g *groupServiceImpl) GenerateQRCode(ctx context.Context, id string) (file []byte, err error) {
	if _, repoErr := g.groupRepository.FindByID(ctx, id); repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	file, genErr := g.qrCodeGenerator.GenerateQRCode(id, qrcode.Medium, 2048)
	if genErr != nil {
		err = service.MapError(genErr)
	}
	return
}

func mapToModel(e entity.Group) response.Group {
	properties := make([]response.Property, len(e.Properties))

	for i, prop := range e.Properties {
		properties[i].ID = prop.ID
		properties[i].Name = prop.Name
		properties[i].Amount = prop.Amount
		properties[i].Description = prop.Description
	}

	return response.Group{
		ID:     e.ID,
		Name:   e.Name,
		Leader: e.Leader,
		Address: response.Address{
			ID:           e.Address.ID,
			Address:      e.Address.Address,
			VillageID:    e.Address.VillageID,
			VillageName:  e.Address.VillageName,
			DistrictID:   e.Address.DistrictID,
			DistrictName: e.Address.DistrictName,
			RegencyID:    e.Address.RegencyID,
			RegencyName:  e.Address.RegencyName,
			ProvinceID:   e.Address.ProvinceID,
			ProvinceName: e.Address.ProvinceName,
		},
		Properties: properties,
	}
}

func mapToModels(entities []entity.Group) []response.Group {
	groups := make([]response.Group, len(entities))

	for i, e := range entities {
		groups[i] = mapToModel(e)
	}

	return groups
}
