package showschedule

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/repository"
	mgr "github.com/erikrios/reog-apps-apis/repository/group/mocks"
	mssr "github.com/erikrios/reog-apps-apis/repository/showschedule/mocks"
	"github.com/erikrios/reog-apps-apis/service"
	mig "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockShowScheduleRepo := &mssr.ShowScheduleRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}

	var showScheduleService ShowScheduleService = NewShowScheduleServiceImpl(
		mockShowScheduleRepo,
		mockGroupRepo,
		mockIDGen,
	)

	testCases := []struct {
		name                    string
		inputCreateShowSchedule payload.CreateShowSchedule
		expectedID              string
		expectedError           error
		mockBehaviours          func()
	}{
		{
			name: "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "",
				Place:    "Lapangan Bungkal",
				StartOn:  "",
				FinishOn: "",
			},
			expectedID:     "",
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name: "it should return service.ErrTimeParsing error, when StartOn payload is invalid",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "Feb 02 06 15:04 WIB",
				FinishOn: "Feb 02 06 17:05 WIB",
			},
			expectedID:     "",
			expectedError:  service.ErrTimeParsing,
			mockBehaviours: func() {},
		},
		{
			name: "it should return service.ErrTimeParsing error, when FinishOn payload is invalid",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "Feb 02 06 17:05 WIB",
			},
			expectedID:     "",
			expectedError:  service.ErrTimeParsing,
			mockBehaviours: func() {},
		},
		{
			name: "it should return service.ErrDataNotFound error, when group repository return an error",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "02 Feb 06 15:04 WIB",
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
			name: "it should return service.ErrRepository error, when id generator return an error",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "02 Feb 06 15:04 WIB",
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

				mockIDGen.On("GenerateShowScheduleID").Return(
					func() string {
						return ""
					},
					func() error {
						return errors.New("error generate show schedule id")
					},
				).Once()
			},
		},
		{
			name: "it should return service.ErrRepository error, when show schedule repository return an error",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "02 Feb 06 15:04 WIB",
			},
			expectedID:    "s-EuKgD1O",
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

				mockIDGen.On("GenerateShowScheduleID").Return(
					func() string {
						return "s-EuKgD1O"
					},
					func() error {
						return nil
					},
				).Once()

				mockShowScheduleRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.ShowSchedule{})),
				).Return(
					func(ctx context.Context, e entity.ShowSchedule) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name: "it should return a valid ID, when no error is returned",
			inputCreateShowSchedule: payload.CreateShowSchedule{
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "05 May 22 13:00 WIB",
				FinishOn: "05 May 22 17:00 WIB",
			},
			expectedID:    "s-EuKgD1O",
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

				mockIDGen.On("GenerateShowScheduleID").Return(
					func() string {
						return "s-EuKgD1O"
					},
					func() error {
						return nil
					},
				).Once()

				mockShowScheduleRepo.On(
					"Insert",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.ShowSchedule{})),
				).Return(
					func(ctx context.Context, e entity.ShowSchedule) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotID, gotErr := showScheduleService.Create(context.Background(), testCase.inputCreateShowSchedule)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, testCase.expectedID, gotID)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	mockShowScheduleRepo := &mssr.ShowScheduleRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}

	var showScheduleService ShowScheduleService = NewShowScheduleServiceImpl(
		mockShowScheduleRepo,
		mockGroupRepo,
		mockIDGen,
	)

	testCases := []struct {
		name                  string
		expectedShowSchedules []response.ShowSchedule
		expectedError         error
		mockBehaviours        func()
	}{
		{
			name:          "it should return service.ErrRepository error, when show schedule repository return an error",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindAll",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
				).Return(
					func(ctx context.Context) []entity.ShowSchedule {
						return []entity.ShowSchedule{}
					},
					func(ctx context.Context) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return a valid show schedules, when no error is returned",
			expectedError: nil,
			expectedShowSchedules: []response.ShowSchedule{
				{
					ID:       "s-EuKgD1O",
					GroupID:  "g-xyz",
					Place:    "Lapangan Bungkal",
					StartOn:  time.Now().Format(time.RFC822),
					FinishOn: time.Now().Add(3 * time.Hour).Format(time.RFC822),
				},
			},
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindAll",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
				).Return(
					func(ctx context.Context) []entity.ShowSchedule {
						return []entity.ShowSchedule{
							{
								ID:       "s-EuKgD1O",
								GroupID:  "g-xyz",
								Place:    "Lapangan Bungkal",
								StartOn:  time.Now(),
								FinishOn: time.Now().Add(3 * time.Hour),
							},
						}
					},
					func(ctx context.Context) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotShowSchedules, gotErr := showScheduleService.GetAll(context.Background())

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.ElementsMatch(t, testCase.expectedShowSchedules, gotShowSchedules)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	mockShowScheduleRepo := &mssr.ShowScheduleRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}

	var showScheduleService ShowScheduleService = NewShowScheduleServiceImpl(
		mockShowScheduleRepo,
		mockGroupRepo,
		mockIDGen,
	)

	testCases := []struct {
		name                 string
		expectedShowSchedule response.ShowScheduleDetails
		expectedError        error
		mockBehaviours       func()
	}{
		{
			name:          "it should return service.ErrDataNotFound error, when show schedule repository return an error",
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.ShowSchedule {
						return entity.ShowSchedule{}
					},
					func(ctx context.Context, id string) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:          "it should return service.ErrRepository error, when show schedule repository return an error",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.ShowSchedule {
						return entity.ShowSchedule{}
					},
					func(ctx context.Context, id string) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return service.ErrRepository error, when group repository return an error",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.ShowSchedule {
						return entity.ShowSchedule{
							ID:       "s-EuKgD1O",
							GroupID:  "g-xyz",
							Place:    "Lapangan Bungkal",
							StartOn:  time.Now(),
							FinishOn: time.Now().Add(3 * time.Hour),
						}
					},
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()

				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{}
					},
					func(ctx context.Context, id string) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return a valid show schedules, when no error is returned",
			expectedError: nil,
			expectedShowSchedule: response.ShowScheduleDetails{
				ID:        "s-EuKgD1O",
				GroupID:   "g-xyz",
				GroupName: "Paguyuban Reog",
				Place:     "Lapangan Bungkal",
				StartOn:   time.Now().Format(time.RFC822),
				FinishOn:  time.Now().Add(3 * time.Hour).Format(time.RFC822),
			},
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.ShowSchedule {
						return entity.ShowSchedule{
							ID:       "s-EuKgD1O",
							GroupID:  "g-xyz",
							Place:    "Lapangan Bungkal",
							StartOn:  time.Now(),
							FinishOn: time.Now().Add(3 * time.Hour),
						}
					},
					func(ctx context.Context, id string) error {
						return nil
					},
				).Once()

				mockGroupRepo.On(
					"FindByID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) entity.Group {
						return entity.Group{
							ID:   "g-xyz",
							Name: "Paguyuban Reog",
						}
					},
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

			gotShowSchedule, gotErr := showScheduleService.GetByID(context.Background(), "g-xyz")

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, testCase.expectedShowSchedule, gotShowSchedule)
			}
		})
	}
}

func TestGetByGroupID(t *testing.T) {
	mockShowScheduleRepo := &mssr.ShowScheduleRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}

	var showScheduleService ShowScheduleService = NewShowScheduleServiceImpl(
		mockShowScheduleRepo,
		mockGroupRepo,
		mockIDGen,
	)

	testCases := []struct {
		name                  string
		expectedShowSchedules []response.ShowSchedule
		expectedError         error
		mockBehaviours        func()
	}{
		{
			name:          "it should return service.ErrRepository error, when group repository return an error",
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
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return service.ErrRepository error, when show schedule repository return an error",
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

				mockShowScheduleRepo.On(
					"FindByGroupID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) []entity.ShowSchedule {
						return []entity.ShowSchedule{}
					},
					func(ctx context.Context, id string) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:          "it should return a valid show schedules, when no error is returned",
			expectedError: nil,
			expectedShowSchedules: []response.ShowSchedule{
				{
					ID:       "s-EuKgD1O",
					GroupID:  "g-xyz",
					Place:    "Lapangan Bungkal",
					StartOn:  time.Now().Format(time.RFC822),
					FinishOn: time.Now().Add(3 * time.Hour).Format(time.RFC822),
				},
			},
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

				mockShowScheduleRepo.On(
					"FindByGroupID",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
				).Return(
					func(ctx context.Context, id string) []entity.ShowSchedule {
						return []entity.ShowSchedule{
							{
								ID:       "s-EuKgD1O",
								GroupID:  "g-xyz",
								Place:    "Lapangan Bungkal",
								StartOn:  time.Now(),
								FinishOn: time.Now().Add(3 * time.Hour),
							},
						}
					},
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

			gotShowSchedules, gotErr := showScheduleService.GetByGroupID(context.Background(), "g-xyz")

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.ElementsMatch(t, testCase.expectedShowSchedules, gotShowSchedules)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	mockShowScheduleRepo := &mssr.ShowScheduleRepository{}
	mockGroupRepo := &mgr.GroupRepository{}
	mockIDGen := &mig.IDGenerator{}

	var showScheduleService ShowScheduleService = NewShowScheduleServiceImpl(
		mockShowScheduleRepo,
		mockGroupRepo,
		mockIDGen,
	)

	testCases := []struct {
		name                    string
		inputID                 string
		inputUpdateShowSchedule payload.UpdateShowSchedule
		expectedError           error
		mockBehaviours          func()
	}{
		{
			name:    "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "",
				FinishOn: "",
			},
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name:    "it should return service.ErrTimeParsing error, when StartOn payload is invalid",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "Feb 02 06 15:04 WIB",
				FinishOn: "Feb 02 06 17:05 WIB",
			},
			expectedError:  service.ErrTimeParsing,
			mockBehaviours: func() {},
		},
		{
			name:    "it should return service.ErrTimeParsing error, when FinishOn payload is invalid",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "Feb 02 06 17:05 WIB",
			},
			expectedError:  service.ErrTimeParsing,
			mockBehaviours: func() {},
		},
		{
			name:    "it should return service.ErrDataNotFound error, when show schedule repository return an error",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "02 Feb 06 15:04 WIB",
			},
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.ShowSchedule{})),
				).Return(
					func(ctx context.Context, id string, e entity.ShowSchedule) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name:    "it should return service.ErrRepository error, when show schedule repository return an error",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "02 Feb 06 15:04 WIB",
				FinishOn: "02 Feb 06 15:04 WIB",
			},
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.ShowSchedule{})),
				).Return(
					func(ctx context.Context, id string, e entity.ShowSchedule) error {
						return repository.ErrDatabase
					},
				).Once()
			},
		},
		{
			name:    "it should return a nil error, when no error is returned",
			inputID: "s-EuKgD1O",
			inputUpdateShowSchedule: payload.UpdateShowSchedule{
				Place:    "Lapangan Bungkal",
				StartOn:  "05 May 22 13:00 WIB",
				FinishOn: "05 May 22 17:00 WIB",
			},
			expectedError: nil,
			mockBehaviours: func() {
				mockShowScheduleRepo.On(
					"Update",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", "")),
					mock.AnythingOfType(fmt.Sprintf("%T", entity.ShowSchedule{})),
				).Return(
					func(ctx context.Context, id string, e entity.ShowSchedule) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()

			gotErr := showScheduleService.Update(context.Background(), testCase.inputID, testCase.inputUpdateShowSchedule)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
}
