package admin

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository"
	mr "github.com/erikrios/reog-apps-apis/repository/admin/mocks"
	"github.com/erikrios/reog-apps-apis/service"
	mpg "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	mtg "github.com/erikrios/reog-apps-apis/utils/generator/mocks"
	_ "github.com/erikrios/reog-apps-apis/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	mockRepo := &mr.AdminRepository{}
	mockPwdGen := &mpg.PasswordGenerator{}
	mockTknGen := &mtg.TokenGenerator{}

	var adminService AdminService = NewAdminServiceImpl(mockRepo, mockPwdGen, mockTknGen)

	testCases := []struct {
		name            string
		inputCredential payload.Credential
		expectedToken   string
		expectedError   error
		mockBehaviours  func()
	}{
		{
			name: "it should return service.ErrInvalidPayload error, when payload is invalid",
			inputCredential: payload.Credential{
				Username: "a",
				Password: "",
			},
			expectedToken:  "",
			expectedError:  service.ErrInvalidPayload,
			mockBehaviours: func() {},
		},
		{
			name: "it should return service.ErrDataNotFound error, when repository return an error",
			inputCredential: payload.Credential{
				Username: "erikris",
				Password: "secret",
			},
			expectedToken: "",
			expectedError: service.ErrDataNotFound,
			mockBehaviours: func() {
				mockRepo.On("FindByUsername", mock.AnythingOfType(fmt.Sprintf("%T", context.Background())), mock.AnythingOfType("string")).Return(
					func(ctx context.Context, username string) entity.Admin {
						return entity.Admin{}
					},
					func(ctx context.Context, username string) error {
						return repository.ErrRecordNotFound
					},
				).Once()
			},
		},
		{
			name: "it should return service.ErrCredentialNotMatch error, when password generator return an error",
			inputCredential: payload.Credential{
				Username: "erikrios",
				Password: "secret",
			},
			expectedToken: "",
			expectedError: service.ErrCredentialNotMatch,
			mockBehaviours: func() {
				mockRepo.On(
					"FindByUsername",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType("string"),
				).Return(
					func(ctx context.Context, username string) entity.Admin {
						return entity.Admin{
							ID:       "a-xy",
							Username: "erikrios",
							Name:     "Erik Rio Setiawan",
							Password: "secret",
						}
					},
					func(ctx context.Context, username string) error {
						return nil
					},
				).Once()

				mockPwdGen.On(
					"CompareHashAndPassword",
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
				).Return(
					func(hashedPassword []byte, password []byte) error {
						return errors.New("error compare hash and password")
					},
				).Once()
			},
		},
		{
			name: "it should return service.ErrRepository error, when token generator return an error",
			inputCredential: payload.Credential{
				Username: "erikrios",
				Password: "secret",
			},
			expectedToken: "",
			expectedError: service.ErrRepository,
			mockBehaviours: func() {
				mockRepo.On(
					"FindByUsername",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType("string"),
				).Return(
					func(ctx context.Context, username string) entity.Admin {
						return entity.Admin{
							ID:       "a-xy",
							Username: "erikrios",
							Name:     "Erik Rio Setiawan",
							Password: "secret",
						}
					},
					func(ctx context.Context, username string) error {
						return nil
					},
				).Once()

				mockPwdGen.On(
					"CompareHashAndPassword",
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
				).Return(
					func(hashedPassword []byte, password []byte) error {
						return nil
					},
				).Once()
				mockTknGen.On(
					"GenerateToken",
					mock.AnythingOfType("string"),
					mock.AnythingOfType("string"),
				).Return(
					func(id string, username string) string {
						return ""
					},
					func(id string, username string) error {
						return errors.New("error generate token")
					},
				).Once()
			},
		},
		{
			name: "it should return a valid token, when no error is returned",
			inputCredential: payload.Credential{
				Username: "erikrios",
				Password: "secret",
			},
			expectedToken: "generatedtoken",
			expectedError: nil,
			mockBehaviours: func() {
				mockRepo.On(
					"FindByUsername",
					mock.AnythingOfType(fmt.Sprintf("%T", context.Background())),
					mock.AnythingOfType("string"),
				).Return(
					func(ctx context.Context, username string) entity.Admin {
						return entity.Admin{
							ID:       "a-xy",
							Username: "erikrios",
							Name:     "Erik Rio Setiawan",
							Password: "secret",
						}
					},
					func(ctx context.Context, username string) error {
						return nil
					},
				).Once()

				mockPwdGen.On(
					"CompareHashAndPassword",
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
					mock.AnythingOfType(fmt.Sprintf("%T", []byte{})),
				).Return(
					func(hashedPassword []byte, password []byte) error {
						return nil
					},
				).Once()
				mockTknGen.On(
					"GenerateToken",
					mock.AnythingOfType("string"),
					mock.AnythingOfType("string"),
				).Return(
					func(id string, username string) string {
						return "generatedtoken"
					},
					func(id string, username string) error {
						return nil
					},
				).Once()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviours()
			gotToken, gotErr := adminService.Login(context.Background(), testCase.inputCredential)

			if testCase.expectedError != nil {
				assert.ErrorIs(t, gotErr, testCase.expectedError)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, testCase.expectedToken, gotToken)
			}
		})
	}
}
