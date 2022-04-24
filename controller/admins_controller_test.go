package controller

import (
	"testing"

	"github.com/erikrios/reog-apps-apis/service/admin/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	mockService := &mocks.AdminService{}
	controller := NewAdminsController(mockService)
	g := echo.New().Group("/api/v1")
	controller.Route(g)
	assert.NotNil(t, controller)
}
