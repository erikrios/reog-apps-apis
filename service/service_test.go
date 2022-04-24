package service

import (
	"errors"
	"testing"

	"github.com/erikrios/reog-apps-apis/repository"
	"github.com/stretchr/testify/assert"
)

func TestMapError(t *testing.T) {
	testCases := []struct {
		name          string
		inputError    error
		expectedError error
	}{
		{
			name:          "it should return service.ErrDataNotFound, when input error is repository.ErrRecordNotFound",
			inputError:    repository.ErrRecordNotFound,
			expectedError: ErrDataNotFound,
		},
		{
			name:          "it should return service.ErrRepository, when input error is repository.ErrDatabase",
			inputError:    repository.ErrDatabase,
			expectedError: ErrRepository,
		},
		{
			name:          "it should return service.ErrDataAlreadyExists, when input error is repository.ErrRecordAlreadyExists",
			inputError:    repository.ErrRecordAlreadyExists,
			expectedError: ErrDataAlreadyExists,
		},
		{
			name:          "it should return service.ErrRepository, when input error is general error",
			inputError:    errors.New("error general"),
			expectedError: ErrRepository,
		},
	}

	for _, testCase := range testCases {
		gotError := MapError(testCase.inputError)
		assert.ErrorIs(t, gotError, testCase.expectedError)
	}
}
