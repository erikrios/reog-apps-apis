package repository

import "errors"

var (
	ErrEntityNotFound = errors.New("repository: entity with given params not found")
	ErrDatabase       = errors.New("repository: something wrong with the database")
)
