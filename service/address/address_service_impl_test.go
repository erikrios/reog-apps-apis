package address

import (
	"context"
	"fmt"
	"testing"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository"
	mar "github.com/erikrios/reog-apps-apis/repository/address/mocks"
	mvr "github.com/erikrios/reog-apps-apis/repository/village/mocks"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	mockAddressRepo := &mar.AddressRepository{}
	mockVillageRepo := &mvr.VillageRepository{}

	var addressService AddressService = NewAddressServiceImpl(
		mockAddressRepo,
		mockVillageRepo,
	)

	testCases := []struct {
		name               string
		inputID            string
		inputUpdateAddress payload.UpdateAddress
		expectedError      error
		mockBehaviours     func()
	}{
		{
			name:    "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputID: "g-xyz",
			inputUpdateAddress: payload.UpdateAddress{
				Address:   "R",
				VillageID: "3502030007",
			},
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name:    "it should return service.ErrDataNotFound error, when village repository return an error",
			inputID: "g-xyz",
			inputUpdateAddress: payload.UpdateAddress{
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502031117",
			},
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockVillageRepo.On("FindByID", mock.AnythingOfType("string")).Return(
					func(id string) entity.Village {
						return entity.Village{}
					},
					func(id string) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:    "it should return service.ErrRepository error, when addresss repository return an error",
			inputID: "g-xyz",
			inputUpdateAddress: payload.UpdateAddress{
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockVillageRepo.On("FindByID", mock.AnythingOfType("string")).Return(
					func(id string) entity.Village {
						return entity.Village{}
					},
					func(id string) error {
						return nil
					},
				).Once()

				mockAddressRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Address{})),
				).Return(
					func(ctx context.Context, id string, group entity.Address) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:    "it should return service.ErrDataNotFound error, when addresss repository return an error",
			inputID: "g-xyz",
			inputUpdateAddress: payload.UpdateAddress{
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockVillageRepo.On("FindByID", mock.AnythingOfType("string")).Return(
					func(id string) entity.Village {
						return entity.Village{}
					},
					func(id string) error {
						return nil
					},
				).Once()

				mockAddressRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Address{})),
				).Return(
					func(ctx context.Context, id string, group entity.Address) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:    "it should return nil error, when no error is returned",
			inputID: "g-xyz",
			inputUpdateAddress: payload.UpdateAddress{
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedError: nil,
			mockBehaviours: func() {
				mockVillageRepo.On("FindByID", mock.AnythingOfType("string")).Return(
					func(id string) entity.Village {
						return entity.Village{}
					},
					func(id string) error {
						return nil
					},
				).Once()

				mockAddressRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Address{})),
				).Return(
					func(ctx context.Context, id string, group entity.Address) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()
			gotErr := addressService.Update(context.Background(), testCase.inputID, testCase.inputUpdateAddress)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
