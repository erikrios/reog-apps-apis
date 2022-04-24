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
	"github.com/skip2/go-qrcode"
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

func TestUpdate(t *testing.T) {
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
		inputID             string
		inputUpdateProperty payload.UpdateProperty
		expectedError       error
		mockBehaviours      func()
	}{
		{
			name:    "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputID: "p-Gx9LkMn",
			inputUpdateProperty: payload.UpdateProperty{
				Name:        "D",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name:    "it should return service.ErrDataNotFound error, when property repository return an error",
			inputID: "p-Gx9LkMn",
			inputUpdateProperty: payload.UpdateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Property{})),
				).Return(
					func(ctx context.Context, id string, p entity.Property) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:    "it should return service.ErrRepository error, when property repository return an error",
			inputID: "p-Gx9LkMn",
			inputUpdateProperty: payload.UpdateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Property{})),
				).Return(
					func(ctx context.Context, id string, p entity.Property) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:    "it should return nil error, when no error is returned",
			inputID: "p-Gx9LkMn",
			inputUpdateProperty: payload.UpdateProperty{
				Name:        "Dadak Merak",
				Description: "Ini Deskripsi Dadak Merak",
				Amount:      1,
			},
			expectedError: nil,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Property{})),
				).Return(
					func(ctx context.Context, id string, p entity.Property) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotErr := propertyService.Update(context.Background(), testCase.inputID, testCase.inputUpdateProperty)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
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
		name           string
		inputID        string
		expectedError  error
		mockBehaviours func()
	}{
		{
			name:          "it should return service.ErrDataNotFound error, when property repository return an error",
			inputID:       "p-Gx9LkMn",
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Delete",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:          "it should return service.ErrRepository error, when property repository return an error",
			inputID:       "p-Gx9LkMn",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Delete",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return nil error, when no error is returned",
			inputID:       "p-Gx9LkMn",
			expectedError: nil,
			mockBehaviours: func() {
				mockPropertyRepo.On(
					"Delete",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotErr := propertyService.Delete(context.Background(), testCase.inputID)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestGenerateQRCode(t *testing.T) {
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
		name           string
		inputID        string
		expectedFile   []byte
		expectedError  error
		mockBehaviours func()
	}{
		{
			name:          "it should return service.ErrRepository error, when QR Code Generator return an error",
			inputID:       "g-xyz",
			expectedFile:  []byte{},
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockQRGen.On(
					"GenerateQRCode",
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", qrcode.Medium)),
					mock.AnythingOfType(fmt.Sprintf("%T", 2048)),
				).Return(
					func(id string, level qrcode.RecoveryLevel, size int) []byte {
						return []byte{}
					},
					func(id string, level qrcode.RecoveryLevel, size int) error {
						return errors.New("error generate qrcode")
					},
				).Once()
			},
		},
		{
			name:          "it should return a valid file, when no error is returned",
			inputID:       "g-xyz",
			expectedFile:  []byte{1},
			expectedError: nil,
			mockBehaviours: func() {
				mockQRGen.On(
					"GenerateQRCode",
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", qrcode.Medium)),
					mock.AnythingOfType(fmt.Sprintf("%T", 2048)),
				).Return(
					func(id string, level qrcode.RecoveryLevel, size int) []byte {
						return []byte{1}
					},
					func(id string, level qrcode.RecoveryLevel, size int) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotFile, gotErr := propertyService.GenerateQRCode(context.Background(), testCase.inputID)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.ElementsMatch(t, testCase.expectedFile, gotFile)
			}
		})
	}
}
