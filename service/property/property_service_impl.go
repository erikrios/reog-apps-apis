package property

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/property"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"gopkg.in/validator.v2"
)

type propertyServiceImpl struct {
	propertyRepository property.PropertyRepository
	groupRepository    group.GroupRepository
	idGenerator        generator.IDGenerator
	qrCodeGenerator    generator.QRCodeGenerator
}

func NewPropertyServiceImpl(
	propertyRepository property.PropertyRepository,
	groupRepository group.GroupRepository,
	idGenerator generator.IDGenerator,
	qrCodeGenerator generator.QRCodeGenerator,
) *propertyServiceImpl {
	return &propertyServiceImpl{
		propertyRepository: propertyRepository,
		groupRepository:    groupRepository,
		idGenerator:        idGenerator,
		qrCodeGenerator:    qrCodeGenerator,
	}
}

func (p *propertyServiceImpl) Create(ctx context.Context, groupID string, payload payload.CreateProperty) (id string, err error) {
	if validateErr := validator.Validate(payload); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	if _, repoErr := p.groupRepository.FindByID(ctx, groupID); repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	id, genErr := p.idGenerator.GeneratePropertyID()
	if genErr != nil {
		err = service.MapError(genErr)
		return
	}

	property := entity.Property{
		ID:          id,
		Name:        payload.Name,
		Description: payload.Description,
		Amount:      payload.Amount,
		GroupID:     groupID,
	}

	if repoErr := p.propertyRepository.Insert(ctx, property); repoErr != nil {
		err = service.MapError(repoErr)
	}

	return
}

func (p *propertyServiceImpl) Update(ctx context.Context, id string, payload payload.UpdateProperty) (err error) {
	if validateErr := validator.Validate(payload); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	property := entity.Property{
		ID:          id,
		Name:        payload.Name,
		Description: payload.Description,
		Amount:      payload.Amount,
	}

	if repoErr := p.propertyRepository.Update(ctx, id, property); repoErr != nil {
		err = service.MapError(repoErr)
	}
	return
}

func (p *propertyServiceImpl) Delete(ctx context.Context, id string) (err error) {
	if repoErr := p.propertyRepository.Delete(ctx, id); repoErr != nil {
		err = service.MapError(repoErr)
	}
	return
}

func (p *propertyServiceImpl) GeterateQRCode(ctx context.Context, id string) (file []byte, err error) {
	return
}
