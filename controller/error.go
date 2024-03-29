package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/erikrios/reog-apps-apis/service"
	"github.com/labstack/echo/v4"
)

func newErrorResponse(err error) *echo.HTTPError {
	var statusCode int
	var message string

	if errors.Is(err, service.ErrDataNotFound) {
		statusCode = http.StatusNotFound
		message = "Resource with given ID not found."
	} else if errors.Is(err, service.ErrDataAlreadyExists) {
		statusCode = http.StatusBadRequest
		message = "Data already exists."
	} else if errors.Is(err, service.ErrTimeParsing) {
		statusCode = http.StatusBadRequest
		message = "Invalid time format. Please use RFC822 time format (02 Jan 06 15:04 MST)"
	} else if errors.Is(err, service.ErrInvalidPayload) {
		statusCode = http.StatusBadRequest
		message = "Invalid payload. Please check the payload schema in the API Documentation."
	} else if errors.Is(err, service.ErrCredentialNotMatch) {
		statusCode = http.StatusUnauthorized
		message = "Username and password not match."
	} else if errors.Is(err, service.ErrRepository) {
		statusCode = http.StatusInternalServerError
		message = "Something went wrong."
	} else {
		statusCode = http.StatusInternalServerError
		message = "Unknown Error."
		log.Println(err)
	}

	return echo.NewHTTPError(statusCode, message)
}
