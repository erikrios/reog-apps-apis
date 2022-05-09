package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/showschedule/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRouteShowSchedules(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}
	controller := NewShowSchedulesController(mockShowScheduleService)
	g := echo.New().Group("/api/v1")
	controller.Route(g)
	assert.NotNil(t, controller)
}

func TestPostCreateShowSchedule(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.CreateShowSchedule{
			GroupID:  "g-xyz",
			Place:    "Lapangan Bungkal",
			StartOn:  "09 May 22 13:00 WIB",
			FinishOn: "09 May 22 17:00 WIB",
		}

		dummyID := "p-aBcdEfL"
		dummyIDResponse := map[string]any{"id": dummyID}
		dummyResp := model.NewResponse("success", "show schedule successfully created", dummyIDResponse)

		mockShowScheduleService.On(
			"Create",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateShowSchedule{})),
		).Return(
			func(ctx context.Context, p payload.CreateShowSchedule) string {
				return dummyID
			},
			func(ctx context.Context, p payload.CreateShowSchedule) error {
				return nil
			},
		).Once()

		t.Run("it should return 201 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/shows", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, controller.postCreateShowSchedule(c)) {
				assert.Equal(t, http.StatusCreated, rec.Code)

				body := rec.Body.String()

				gotResponse := make(map[string]any)

				if err := json.Unmarshal([]byte(body), &gotResponse); assert.NoError(t, err) {
					gotID := gotResponse["data"].(map[string]any)["id"].(string)
					assert.Equal(t, dummyResp.Data["id"], gotID)
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		dummyReq := payload.CreateShowSchedule{
			GroupID:  "g-xyz",
			Place:    "Lapangan Bungkal",
			StartOn:  "09 May 22 13:00 WIB",
			FinishOn: "09 May 22 17:00 WIB",
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.CreateShowSchedule
			expectedStatusCode   int
			expectedErrorMessage string
			mockBehaviour        func()
		}{
			{
				name:                 "it should return 400 status code, when payload is invalid",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusBadRequest,
				expectedErrorMessage: "Invalid payload. Please check the payload schema in the API Documentation.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateShowSchedule{})),
					).Return(
						func(ctx context.Context, p payload.CreateShowSchedule) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateShowSchedule) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when Group ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateShowSchedule{})),
					).Return(
						func(ctx context.Context, p payload.CreateShowSchedule) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateShowSchedule) error {
							return service.ErrDataNotFound
						},
					).Once()
				},
			},
			{
				name:                 "it should return 500 status code, when error happened",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusInternalServerError,
				expectedErrorMessage: "Something went wrong.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateShowSchedule{})),
					).Return(
						func(ctx context.Context, p payload.CreateShowSchedule) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateShowSchedule) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewShowSchedulesController(mockShowScheduleService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/api/v1/shows", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.postCreateShowSchedule(c)
				if assert.Error(t, gotError) {
					if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
						assert.Equal(t, testCase.expectedStatusCode, echoHTTPError.Code)
						assert.Equal(t, testCase.expectedErrorMessage, echoHTTPError.Message)
					}
				}
			})
		}
	})
}
