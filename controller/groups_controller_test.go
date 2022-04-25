package controller

import (
	"testing"

	mas "github.com/erikrios/reog-apps-apis/service/address/mocks"
	mgs "github.com/erikrios/reog-apps-apis/service/group/mocks"
	mps "github.com/erikrios/reog-apps-apis/service/property/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
