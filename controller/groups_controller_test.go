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
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/groups")

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
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/api/v1/groups")

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
