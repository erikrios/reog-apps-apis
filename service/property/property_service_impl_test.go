package property

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository"
	mgr "github.com/erikrios/reog-apps-apis/repository/group/mocks"
	mpr "github.com/erikrios/reog-apps-apis/repository/property/mocks"
	"github.com/erikrios/reog-apps-apis/service"
	mig "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	mqg "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockPropertyRepo := &mpr.PropertyRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}
	mockQRGen := &mqg.QRCodeGenerator{}

	var propertyService PropertyService = NewPropertyServiceImpl(
		mockPropertyRepo,
		mockGroupRepo,
		mockIDGen,
		mockQRGen,
	)

	testCases := []struct {
		name                string
		inputGroupID        string
		inputCreateProperty payload.CreateProperty
		expectedID          string
		expectedError       error
		mockBehaviours      func()
	}{
		{
			name:         "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputGroupID: "g-xyz",
			inputCreateProperty: payload.CreateProperty{
				Name:        "D",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedID:     "",
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name:         "it should return service.ErrDataNotFound error, when group repository return an error",
			inputGroupID: "g-xyz",
			inputCreateProperty: payload.CreateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedID:    "",
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{}
					},
					func(ctx context.Context, id string) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:         "it should return service.ErrRepository error, when id generator return an error",
			inputGroupID: "g-xyz",
			inputCreateProperty: payload.CreateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedID:    "",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{}
					},
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()

				mockIDGen.On("GeneratePropertyID").Return(
					func() string {
						return ""
					},
					func() error {
						return errors.New("error generate group id")
					},
				).Once()
			},
		},
		{
			name:         "it should return service.ErrRepository error, when propety repository return an error",
			inputGroupID: "g-xyz",
			inputCreateProperty: payload.CreateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedID:    "",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{}
					},
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()

				mockIDGen.On("GeneratePropertyID").Return(
					func() string {
						return "p-Gx9LkMn"
					},
					func() error {
						return nil
					},
				).Once()

				mockPropertyRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Property{})),
				).Return(
					func(ctx context.Context, p entity.Property) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:         "it should return a valid ID, when no error is returned",
			inputGroupID: "g-xyz",
			inputCreateProperty: payload.CreateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedID:    "p-Gx9LkMn",
			expectedError: nil,
			mockBehaviours: func() {
				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{}
					},
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()

				mockIDGen.On("GeneratePropertyID").Return(
					func() string {
						return "p-Gx9LkMn"
					},
					func() error {
						return nil
					},
				).Once()

				mockPropertyRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Property{})),
				).Return(
					func(ctx context.Context, p entity.Property) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()
			gotID, gotErr := propertyService.Create(context.Background(), testCase.inputGroupID, testCase.inputCreateProperty)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, testCase.expectedID, gotID)
			}
		})
	}
}
