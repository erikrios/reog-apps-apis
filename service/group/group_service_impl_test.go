package group

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository"
	mgr "github.com/erikrios/reog-apps-apis/repository/group/mocks"
	mvr "github.com/erikrios/reog-apps-apis/repository/village/mocks"
	"github.com/erikrios/reog-apps-apis/service"
	mig "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	mqg "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	_ "github.com/erikrios/reog-apps-apis/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockGroupRepo := &mgr.GroupRepository{}
	mockVillageRepo := &mvr.VillageRepository{}
	mockIDGen := &mig.IDGenerator{}
	mockQRGen := &mqg.QRCodeGenerator{}

	var groupService GroupService = NewGroupServiceImpl(
		mockGroupRepo,
		mockVillageRepo,
		mockIDGen,
		mockQRGen,
	)

	testCases := []struct {
		name             string
		inputCreateGroup payload.CreateGroup
		expectedID       string
		expectedError    error
		mockBehaviours   func()
	}{
		{
			name: "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputCreateGroup: payload.CreateGroup{
				Name:      "Paguyuban Reog",
				Leader:    "E",
				Address:   "RT 01 RW 01",
				VillageID: "",
			},
			expectedID:     "",
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name: "it should return service.ErrDataNotFound error, when repository return an error",
			inputCreateGroup: payload.CreateGroup{
				Name:      "Paguyuban Reog",
				Leader:    "Erik R",
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502031117",
			},
			expectedID:    "",
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
			name: "it should return service.ErrRepository error, when id generator return an error",
			inputCreateGroup: payload.CreateGroup{
				Name:      "Paguyuban Reog",
				Leader:    "Erik R",
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedID:    "",
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

				mockIDGen.On("GenerateGroupID").Return(
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
			name: "it should return service.ErrRepository error, when group repository return an error",
			inputCreateGroup: payload.CreateGroup{
				Name:      "Paguyuban Reog",
				Leader:    "Erik R",
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedID:    "",
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
				mockIDGen.On("GenerateGroupID").Return(
					func() string {
						return ""
					},
					func() error {
						return nil
					},
				).Once()

				mockGroupRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Group{})),
				).Return(
					func(ctx context.Context, group entity.Group) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name: "it should return a valid ID, when no error is returned",
			inputCreateGroup: payload.CreateGroup{
				Name:      "Paguyuban Reog",
				Leader:    "Erik R",
				Address:   "RT 01 RW 01 Dukuh Bibis",
				VillageID: "3502030007",
			},
			expectedID:    "g-xyz",
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
				mockIDGen.On("GenerateGroupID").Return(
					func() string {
						return "g-xyz"
					},
					func() error {
						return nil
					},
				).Once()

				mockGroupRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.Group{})),
				).Return(
					func(ctx context.Context, group entity.Group) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()
			gotID, gotErr := groupService.Create(context.Background(), testCase.inputCreateGroup)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, testCase.expectedID, gotID)
			}
		})
	}
}
