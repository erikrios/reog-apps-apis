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
	"github.com/erikrios/reog-apps-apis/service/admin/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoute(t *testing.T) {
	mockService := &mocks.AdminService{}
	controller := NewAdminsController(mockService)
	g := echo.New().Group("/api/v1")
	controller.Route(g)
	assert.NotNil(t, controller)
}

func TestPostLogin(t *testing.T) {
	mockService := &mocks.AdminService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := payload.Credential{
			Username: "admin",
			Password: "secret",
		}

		dummyToken := "generatedtoken"
		dummyTokenResponse := map[string]any{"token": dummyToken}
		dummyResp := model.NewResponse("success", "login successful", dummyTokenResponse)

		mockService.On(
			"Login",
			mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
			mock.AnythingOfType(fmt.Sprintf("%T", payload.Credential{})),
		).Return(
			func(ctx context.Context, p payload.Credential) string {
				return dummyToken
			},
			func(ctx context.Context, p payload.Credential) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response, when there is no error", func(t *testing.T) {
			controller := NewAdminsController(mockService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/admins")

			if assert.NoError(t, controller.postLogin(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				body := rec.Body.String()

				gotResponse := make(map[string]any)

				if err := json.Unmarshal([]byte(body), &gotResponse); assert.NoError(t, err) {
					gotToken := gotResponse["data"].(map[string]any)["token"].(string)
					assert.Equal(t, dummyResp.Data["token"], gotToken)
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {})
	dummyReq := payload.Credential{
		Username: "admin",
		Password: "secret",
	}

	testCases := []struct {
		name                 string
		inputPayload         payload.Credential
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
				mockService.On(
					"Login",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", payload.Credential{})),
				).Return(
					func(ctx context.Context, p payload.Credential) string {
						return ""
					},
					func(ctx context.Context, p payload.Credential) error {
						return service.ErrInvalidPayload
					},
				).Once()
			},
		},
		{
			name:                 "it should return 404 status code, when username not found",
			inputPayload:         dummyReq,
			expectedStatusCode:   http.StatusNotFound,
			expectedErrorMessage: "Resource with given ID not found.",
			mockBehaviour: func() {
				mockService.On(
					"Login",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", payload.Credential{})),
				).Return(
					func(ctx context.Context, p payload.Credential) string {
						return ""
					},
					func(ctx context.Context, p payload.Credential) error {
						return service.ErrDataNotFound
					},
				).Once()
			},
		},
		{
			name:                 "it should return 401 status code, when username and password not match",
			inputPayload:         dummyReq,
			expectedStatusCode:   http.StatusUnauthorized,
			expectedErrorMessage: "Username and password not match.",
			mockBehaviour: func() {
				mockService.On(
					"Login",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", payload.Credential{})),
				).Return(
					func(ctx context.Context, p payload.Credential) string {
						return ""
					},
					func(ctx context.Context, p payload.Credential) error {
						return service.ErrCredentialNotMatch
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
				mockService.On(
					"Login",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType(fmt.Sprintf("%T", payload.Credential{})),
				).Return(
					func(ctx context.Context, p payload.Credential) string {
						return ""
					},
					func(ctx context.Context, p payload.Credential) error {
						return service.ErrRepository
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour()

			controller := NewAdminsController(mockService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/admins")

			gotError := controller.postLogin(c)
			if assert.Error(t, gotError) {
				if echoHTTPError, ok := gotError.(*echo.HTTPError); assert.Equal(t, true, ok) {
					assert.Equal(t, testCase.expectedStatusCode, echoHTTPError.Code)
					assert.Equal(t, testCase.expectedErrorMessage, echoHTTPError.Message)
				}
			}
		})
	}
}
