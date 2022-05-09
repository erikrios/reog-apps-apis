package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
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

func TestGetShowSchedules(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyShowSchedules := []response.ShowSchedule{
			{
				ID:       "s-abcdefg",
				GroupID:  "g-xyz",
				Place:    "Lapangan Bungkal",
				StartOn:  "09 May 22 13:00 WIB",
				FinishOn: "09 May 22 17:00 WIB",
			},
		}

		dummyGroupsResponse := map[string]any{"shows": dummyShowSchedules}
		dummyResp := model.NewResponse("success", "successfully get show schedules", dummyGroupsResponse)

		mockShowScheduleService.On(
			"GetAll",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
		).Return(
			func(ctx context.Context) []response.ShowSchedule {
				return dummyShowSchedules
			},
			func(ctx context.Context) error {
				return nil
			},
		).Once()

		mockShowScheduleService.On(
			"GetByGroupID",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, groupID string) []response.ShowSchedule {
				return dummyShowSchedules
			},
			func(ctx context.Context, groupID string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error and group_id query is empty", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/shows", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, controller.getShowSchedules(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				body := rec.Body.String()

				gotResponse := &model.Response[map[string]any]{}

				if err := json.Unmarshal([]byte(body), &gotResponse); assert.NoError(t, err) {
					reflect.DeepEqual(dummyResp.Data["groups"], gotResponse.Data["groups"])
				}
			}
		})

		t.Run("it should return 200 status code with valid response, when there is no error and group_id query is exists", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/shows?group_id=g-xyz", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, controller.getShowSchedules(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				body := rec.Body.String()

				gotResponse := &model.Response[map[string]any]{}

				if err := json.Unmarshal([]byte(body), &gotResponse); assert.NoError(t, err) {
					reflect.DeepEqual(dummyResp.Data["groups"], gotResponse.Data["groups"])
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		testCases := []struct {
			name                 string
			expectedStatusCode   int
			expectedErrorMessage string
			mockBehaviour        func()
		}{
			{
				name:                 "it should return 500 status code, when error happened",
				expectedStatusCode:   http.StatusInternalServerError,
				expectedErrorMessage: "Something went wrong.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"GetAll",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					).Return(
						func(ctx context.Context) []response.ShowSchedule {
							return []response.ShowSchedule{}
						},
						func(ctx context.Context) error {
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

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/shows", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.getShowSchedules(c)
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

func TestGetShowScheduleByID(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyShow := response.ShowScheduleDetails{
			ID:        "s-abcdefg",
			GroupID:   "g-xyz",
			GroupName: "Paguyuban Reog",
			Place:     "Lapangan Bungkal",
			StartOn:   "09 May 22 13:00 WIB",
			FinishOn:  "09 May 22 17:00 WIB",
		}

		dummyShowsResponse := map[string]any{"show": dummyShow}
		dummyResp := model.NewResponse("success", "successfully get show schedule with id "+dummyShow.ID, dummyShowsResponse)

		mockShowScheduleService.On(
			"GetByID",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) response.ShowScheduleDetails {
				return dummyShow
			},
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/shows", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(dummyShow.ID)

			if assert.NoError(t, controller.getShowScheduleByID(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				body := rec.Body.String()

				gotResponse := &model.Response[map[string]any]{}

				if err := json.Unmarshal([]byte(body), &gotResponse); assert.NoError(t, err) {
					reflect.DeepEqual(dummyResp, gotResponse)
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		testCases := []struct {
			name                 string
			expectedStatusCode   int
			expectedErrorMessage string
			mockBehaviour        func()
		}{
			{
				name:                 "it should return 404 status code, when group ID not found",
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"GetByID",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) response.ShowScheduleDetails {
							return response.ShowScheduleDetails{}
						},
						func(ctx context.Context, id string) error {
							return service.ErrDataNotFound
						},
					).Once()
				},
			},
			{
				name:                 "it should return 500 status code, when error happened",
				expectedStatusCode:   http.StatusInternalServerError,
				expectedErrorMessage: "Something went wrong.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"GetByID",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) response.ShowScheduleDetails {
							return response.ShowScheduleDetails{}
						},
						func(ctx context.Context, id string) error {
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

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/shows", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("s-abcdefg")

				gotError := controller.getShowScheduleByID(c)
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

func TestPutUpdateShowScheduleByID(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.UpdateShowSchedule{
			Place:    "Lapangan Pager",
			StartOn:  "09 May 22 13:00 WIB",
			FinishOn: "09 May 22 17:00 WIB",
		}

		mockShowScheduleService.On(
			"Update",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateShowSchedule{})),
		).Return(
			func(ctx context.Context, id string, p payload.UpdateShowSchedule) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/shows", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("s-abcdefg")

			if assert.NoError(t, controller.putUpdateShowScheduleByID(c)) {
				assert.Equal(t, http.StatusNoContent, rec.Code)
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		dummyReq := payload.UpdateShowSchedule{
			Place:    "Lapangan Pager",
			StartOn:  "09 May 22 13:00 WIB",
			FinishOn: "09 May 22 17:00 WIB",
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.UpdateShowSchedule
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
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateShowSchedule{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateShowSchedule) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when show schedule ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateShowSchedule{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateShowSchedule) error {
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
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateShowSchedule{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateShowSchedule) error {
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
				req := httptest.NewRequest(http.MethodPut, "/api/v1/shows", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("s-abcdefg")

				gotError := controller.putUpdateShowScheduleByID(c)
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

func TestDeleteShowScheduleByID(t *testing.T) {
	mockShowScheduleService := &mocks.ShowScheduleService{}

	t.Run("success scenario", func(t *testing.T) {
		mockShowScheduleService.On(
			"Delete",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewShowSchedulesController(mockShowScheduleService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/shows", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("s-abcdefg")

			if assert.NoError(t, controller.deleteShowScheduleByID(c)) {
				assert.Equal(t, http.StatusNoContent, rec.Code)
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		testCases := []struct {
			name                 string
			expectedStatusCode   int
			expectedErrorMessage string
			mockBehaviour        func()
		}{
			{
				name:                 "it should return 404 status code, when village ID not found",
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Delete",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) error {
							return service.ErrDataNotFound
						},
					).Once()
				},
			},
			{
				name:                 "it should return 500 status code, when error happened",
				expectedStatusCode:   http.StatusInternalServerError,
				expectedErrorMessage: "Something went wrong.",
				mockBehaviour: func() {
					mockShowScheduleService.On(
						"Delete",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) error {
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

				e := echo.New()
				req := httptest.NewRequest(http.MethodDelete, "/api/v1/shows", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("s-abcdefg")

				gotError := controller.deleteShowScheduleByID(c)
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
