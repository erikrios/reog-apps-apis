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
	mas "github.com/erikrios/reog-apps-apis/service/address/mocks"
	mgs "github.com/erikrios/reog-apps-apis/service/group/mocks"
	mps "github.com/erikrios/reog-apps-apis/service/property/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRouteGroups(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}
	controller := NewGroupsController(
		mockGroupService,
		mockPropertyService,
		mockAddressService,
	)
	g := echo.New().Group("/api/v1")
	controller.Route(g)
	assert.NotNil(t, controller)
}

func TestPostCreateGroup(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.CreateGroup{
			Name:      "Paguyuban Reog",
			Leader:    "Erik Rio S",
			Address:   "RT 01 RW 01 Dukuh Bibis",
			VillageID: "3502000",
		}

		dummyID := "g-xyz"
		dummyIDResponse := map[string]any{"id": dummyID}
		dummyResp := model.NewResponse("success", "group successfully created", dummyIDResponse)

		mockGroupService.On(
			"Create",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateGroup{})),
		).Return(
			func(ctx context.Context, p payload.CreateGroup) string {
				return dummyID
			},
			func(ctx context.Context, p payload.CreateGroup) error {
				return nil
			},
		).Once()

		t.Run("it should return 201 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, controller.postCreateGroup(c)) {
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
		dummyReq := payload.CreateGroup{
			Name:      "Paguyuban Reog",
			Leader:    "Erik Rio S",
			Address:   "RT 01 RW 01 Dukuh Bibis",
			VillageID: "3502000",
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.CreateGroup
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
					mockGroupService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateGroup{})),
					).Return(
						func(ctx context.Context, p payload.CreateGroup) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateGroup) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when village ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockGroupService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateGroup{})),
					).Return(
						func(ctx context.Context, p payload.CreateGroup) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateGroup) error {
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
					mockGroupService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateGroup{})),
					).Return(
						func(ctx context.Context, p payload.CreateGroup) string {
							return ""
						},
						func(ctx context.Context, p payload.CreateGroup) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.postCreateGroup(c)
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

func TestGetGroups(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyGroups := []response.Group{
			{
				ID:     "g-xyz",
				Name:   "Paguyuban Reog",
				Leader: "Erik Rio S",
				Address: response.Address{
					ID:           "g-xyz",
					Address:      "RT 01 RW 01 Dukuh Bibis",
					VillageID:    "350211189",
					VillageName:  "Pager",
					DistrictID:   "350211",
					DistrictName: "Bungkal",
					RegencyID:    "3502",
					RegencyName:  "Kabupaten Ponorogo",
					ProvinceID:   "35",
					ProvinceName: "Jawa Timur",
				},
				Properties: []response.Property{
					{
						ID:          "p-Ay8LmNI",
						Name:        "Dadak Merak",
						Description: "Ini adalah deskripsi dadak merak",
						Amount:      1,
					},
				},
			},
		}

		dummyGroupsResponse := map[string]any{"groups": dummyGroups}
		dummyResp := model.NewResponse("success", "successfully get groups", dummyGroupsResponse)

		mockGroupService.On(
			"GetAll",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
		).Return(
			func(ctx context.Context) []response.Group {
				return dummyGroups
			},
			func(ctx context.Context) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, controller.getGroups(c)) {
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
					mockGroupService.On(
						"GetAll",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					).Return(
						func(ctx context.Context) []response.Group {
							return []response.Group{}
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				gotError := controller.getGroups(c)
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

func TestGetGroupByID(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyGroup := response.Group{
			ID:     "g-xyz",
			Name:   "Paguyuban Reog",
			Leader: "Erik Rio S",
			Address: response.Address{
				ID:           "g-xyz",
				Address:      "RT 01 RW 01 Dukuh Bibis",
				VillageID:    "350211189",
				VillageName:  "Pager",
				DistrictID:   "350211",
				DistrictName: "Bungkal",
				RegencyID:    "3502",
				RegencyName:  "Kabupaten Ponorogo",
				ProvinceID:   "35",
				ProvinceName: "Jawa Timur",
			},
			Properties: []response.Property{
				{
					ID:          "p-Ay8LmNI",
					Name:        "Dadak Merak",
					Description: "Ini adalah deskripsi dadak merak",
					Amount:      1,
				},
			},
		}

		dummyGroupsResponse := map[string]any{"group": dummyGroup}
		dummyResp := model.NewResponse("success", "successfully get group with id "+dummyGroup.ID, dummyGroupsResponse)

		mockGroupService.On(
			"GetByID",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) response.Group {
				return dummyGroup
			},
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(dummyGroup.ID)

			if assert.NoError(t, controller.getGroupByID(c)) {
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
					mockGroupService.On(
						"GetByID",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) response.Group {
							return response.Group{}
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
					mockGroupService.On(
						"GetByID",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) response.Group {
							return response.Group{}
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.getGroupByID(c)
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

func TestPutUpdateGroupByID(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.UpdateGroup{
			Name:   "Paguyuban Reog",
			Leader: "Erik Rio S",
		}

		mockGroupService.On(
			"Update",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateGroup{})),
		).Return(
			func(ctx context.Context, id string, p payload.UpdateGroup) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/groups", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")

			if assert.NoError(t, controller.putUpdateGroupByID(c)) {
				assert.Equal(t, http.StatusNoContent, rec.Code)
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		dummyReq := payload.UpdateGroup{
			Name:   "Paguyuban Reog",
			Leader: "Erik Rio S",
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.UpdateGroup
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
					mockGroupService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateGroup{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateGroup) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when village ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockGroupService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateGroup{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateGroup) error {
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
					mockGroupService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateGroup{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateGroup) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPut, "/api/v1/groups", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.putUpdateGroupByID(c)
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

func TestDeleteGroupByID(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		mockGroupService.On(
			"Delete",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")

			if assert.NoError(t, controller.deleteGroupByID(c)) {
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
					mockGroupService.On(
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
					mockGroupService.On(
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodDelete, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.deleteGroupByID(c)
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

func TestGetGenerateQRCode(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		mockGroupService.On(
			"GenerateQRCode",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) []byte {
				return []byte{1}
			},
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id/generate")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")

			if assert.NoError(t, controller.getGenerateQRCode(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
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
					mockGroupService.On(
						"GenerateQRCode",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) []byte {
							return []byte{}
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
					mockGroupService.On(
						"GenerateQRCode",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) []byte {
							return []byte{}
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/generate")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.getGenerateQRCode(c)
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

func TestPutUpdateAddress(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.UpdateAddress{
			Address:   "RT 01 RW 01 Dukuh Bibis",
			VillageID: "3502111",
		}

		mockAddressService.On(
			"Update",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateAddress{})),
		).Return(
			func(ctx context.Context, id string, p payload.UpdateAddress) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/groups/addresses", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")

			if assert.NoError(t, controller.putUpdateAddress(c)) {
				assert.Equal(t, http.StatusNoContent, rec.Code)
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		dummyReq := payload.UpdateAddress{
			Address:   "RT 01 RW 01 Dukuh Bibis",
			VillageID: "3502111",
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.UpdateAddress
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
					mockAddressService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateAddress{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateAddress) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when village ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockAddressService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateAddress{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateAddress) error {
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
					mockAddressService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateAddress{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateAddress) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPut, "/api/v1/groups/addresses", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.putUpdateAddress(c)
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

func TestPostCreateProperty(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.CreateProperty{
			Name:        "Dadak Merak",
			Description: "Ini adalah deskripsi dadak merak",
			Amount:      1,
		}

		dummyID := "p-xyzxyza"
		dummyIDResponse := map[string]any{"id": dummyID}
		dummyResp := model.NewResponse("success", "property successfully created", dummyIDResponse)

		mockPropertyService.On(
			"Create",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateProperty{})),
		).Return(
			func(ctx context.Context, groupID string, p payload.CreateProperty) string {
				return dummyID
			},
			func(ctx context.Context, groupID string, p payload.CreateProperty) error {
				return nil
			},
		).Once()

		t.Run("it should return 201 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id/properties")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")

			if assert.NoError(t, controller.postCreateProperty(c)) {
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
		dummyReq := payload.CreateProperty{
			Name:        "Dadak Merak",
			Description: "Ini adalah deskripsi dadak merak",
			Amount:      1,
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.CreateProperty
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
					mockPropertyService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateProperty{})),
					).Return(
						func(ctx context.Context, groupID string, p payload.CreateProperty) string {
							return ""
						},
						func(ctx context.Context, groupID string, p payload.CreateProperty) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when group ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockPropertyService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateProperty{})),
					).Return(
						func(ctx context.Context, groupID string, p payload.CreateProperty) string {
							return ""
						},
						func(ctx context.Context, groupID string, p payload.CreateProperty) error {
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
					mockPropertyService.On(
						"Create",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.CreateProperty{})),
					).Return(
						func(ctx context.Context, groupID string, p payload.CreateProperty) string {
							return ""
						},
						func(ctx context.Context, groupID string, p payload.CreateProperty) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/api/v1/groups", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/properties")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")

				gotError := controller.postCreateProperty(c)
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

func TestPutUpdateProperty(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.UpdateProperty{
			Name:        "Dadak Merak",
			Description: "Ini adalah dadak merak",
			Amount:      1,
		}

		mockPropertyService.On(
			"Update",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateProperty{})),
		).Return(
			func(ctx context.Context, id string, p payload.UpdateProperty) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/groups", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id/properties/:propertyID")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")
			c.SetParamNames("propertyID")
			c.SetParamValues("p-Ay8LmNI")

			if assert.NoError(t, controller.putUpdateProperty(c)) {
				assert.Equal(t, http.StatusNoContent, rec.Code)
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		dummyReq := payload.UpdateProperty{
			Name:        "Dadak Merak",
			Description: "Ini adalah dadak merak",
			Amount:      1,
		}

		testCases := []struct {
			name                 string
			inputPayload         payload.UpdateProperty
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
					mockPropertyService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateProperty{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateProperty) error {
							return service.ErrInvalidPayload
						},
					).Once()
				},
			},
			{
				name:                 "it should return 404 status code, when property ID not found",
				inputPayload:         dummyReq,
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockPropertyService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateProperty{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateProperty) error {
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
					mockPropertyService.On(
						"Update",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
						mock.AnythingOfType(fmt.Sprintf("%T", payload.UpdateProperty{})),
					).Return(
						func(ctx context.Context, id string, p payload.UpdateProperty) error {
							return service.ErrRepository
						},
					).Once()
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.mockBehaviour()

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)
				requestBody, err := json.Marshal(dummyReq)
				assert.NoError(t, err)

				e := echo.New()
				req := httptest.NewRequest(http.MethodPut, "/api/v1/groups", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/properties/:propertyID")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")
				c.SetParamNames("propertyID")
				c.SetParamValues("p-Ay8LmNI")

				gotError := controller.putUpdateProperty(c)
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

func TestDeleteProperty(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		mockPropertyService.On(
			"Delete",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 204 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id/properties/:propertyID")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")
			c.SetParamNames("propertyID")
			c.SetParamValues("p-Ay8LmNI")

			if assert.NoError(t, controller.deleteProperty(c)) {
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
				name:                 "it should return 404 status code, when property ID not found",
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockPropertyService.On(
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
					mockPropertyService.On(
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodDelete, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/properties/:propertyID")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")
				c.SetParamNames("propertyID")
				c.SetParamValues("p-Ay8LmNI")

				gotError := controller.deleteProperty(c)
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

func TestGetGeneratePropertyQRCode(t *testing.T) {
	mockGroupService := &mgs.GroupService{}
	mockPropertyService := &mps.PropertyService{}
	mockAddressService := &mas.AddressService{}

	t.Run("success scenario", func(t *testing.T) {
		mockPropertyService.On(
			"GenerateQRCode",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", "")),
		).Return(
			func(ctx context.Context, id string) []byte {
				return []byte{1}
			},
			func(ctx context.Context, id string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id/properties/:propertyID")
			c.SetParamNames("id")
			c.SetParamValues("g-xyz")
			c.SetParamNames("propertyID")
			c.SetParamValues("p-Ay8LmNI")

			if assert.NoError(t, controller.getGeneratePropertyQRCode(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
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
				name:                 "it should return 404 status code, when property ID not found",
				expectedStatusCode:   http.StatusNotFound,
				expectedErrorMessage: "Resource with given ID not found.",
				mockBehaviour: func() {
					mockPropertyService.On(
						"GenerateQRCode",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) []byte {
							return []byte{}
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
					mockPropertyService.On(
						"GenerateQRCode",
						mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
						mock.AnythingOfType(fmt.Sprintf("%T", "")),
					).Return(
						func(ctx context.Context, id string) []byte {
							return []byte{}
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

				controller := NewGroupsController(mockGroupService, mockPropertyService, mockAddressService)

				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/:id/properties/:propertyID")
				c.SetParamNames("id")
				c.SetParamValues("g-xyz")
				c.SetParamNames("propertyID")
				c.SetParamValues("p-Ay8LmNI")

				gotError := controller.getGeneratePropertyQRCode(c)
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
