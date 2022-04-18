package service

import (
	"errors"

	"github.com/erikrios/reog-apps-apis/repository"
)

var (
	ErrDataNotFound      = errors.New("service: data with given param not found")
	ErrRepository        = errors.New("service: repository error happened")
	ErrDataAlreadyExists = errors.New("service: data already exists")
	ErrInvalidPayload    = errors.New("service: invalid payload")
)

func MapError(from error) error {
	if errors.Is(from, repository.ErrRecordNotFound) {
		return ErrDataNotFound
	} else if errors.Is(from, repository.ErrDatabase) {
		return ErrRepository
	} else if errors.Is(from, repository.ErrRecordAlreadyExists) {
		return ErrDataAlreadyExists
	} else {
		return ErrRepository
	}
}
