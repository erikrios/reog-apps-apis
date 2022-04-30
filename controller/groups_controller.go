package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/middleware"
	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/address"
	"github.com/erikrios/reog-apps-apis/service/group"
	"github.com/erikrios/reog-apps-apis/service/property"
	"github.com/labstack/echo/v4"
)

type groupsController struct {
	groupService    group.GroupService
	propertyService property.PropertyService
	addressService  address.AddressService
}

func NewGroupsController(
	groupService group.GroupService,
	propertyService property.PropertyService,
	addressService address.AddressService,
) *groupsController {
	return &groupsController{
		groupService:    groupService,
		propertyService: propertyService,
		addressService:  addressService,
	}
}

func (g *groupsController) Route(e *echo.Group) {
	group := e.Group("/groups", middleware.JWTMiddleware())
	group.POST("", g.postCreateGroup)
	group.GET("", g.getGroups)
	group.GET("/:id", g.getGroupByID)
	group.PUT("/:id", g.putUpdateGroupByID)
	group.DELETE("/:id", g.deleteGroupByID)
	group.GET("/:id/generate", g.getGenerateQRCode)
	group.PUT("/addresses/:id", g.putUpdateAddress)
	group.POST("/:id/properties", g.postCreateProperty)
	group.PUT("/:id/properties/:propertyID", g.putUpdateProperty)
	group.DELETE("/:id/properties/:propertyID", g.deleteProperty)
	group.GET("/:id/properties/:propertyID/generate", g.getGeneratePropertyQRCode)
}

// postCreateGroup godoc
// @Summary      Create a Group
// @Description  Create a new group
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default  body  payload.CreateGroup  true  "request body"
// @Security     ApiKeyAuth
// @Success      201  {object}  createGroupResponse
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups [post]
func (g *groupsController) postCreateGroup(c echo.Context) error {
	payload := new(payload.CreateGroup)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	id, err := g.groupService.Create(c.Request().Context(), *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	idResponse := map[string]any{"id": id}
	response := model.NewResponse("success", "group successfully created", idResponse)
	return c.JSON(http.StatusCreated, response)
}

// getGroups     godoc
// @Summary      Get Groups
// @Description  Get Groups
// @Tags         groups
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  groupsResponse
// @Failure      401  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups [get]
func (g *groupsController) getGroups(c echo.Context) error {
	groups, err := g.groupService.GetAll(c.Request().Context())
	if err != nil {
		return newErrorResponse(err)
	}

	groupsResposes := map[string]any{"groups": groups}
	responses := model.NewResponse("success", "successfully get groups", groupsResposes)
	return c.JSON(http.StatusOK, responses)
}

//  getGroupByID godoc
// @Summary      Get Group by ID
// @Description  Get group by ID
// @Tags         groups
// @Produce      json
// @Param        id          path  string  true  "group ID"
// @Security     ApiKeyAuth
// @Success      200  {object}  groupResponse
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [get]
func (g *groupsController) getGroupByID(c echo.Context) error {
	id := c.Param("id")

	group, err := g.groupService.GetByID(c.Request().Context(), id)

	if err != nil {
		return newErrorResponse(err)
	}

	groupResponse := map[string]any{"group": group}
	response := model.NewResponse("success", "successfully get group with id "+id, groupResponse)
	return c.JSON(http.StatusOK, response)
}

// putUpdateGroupByID godoc
// @Summary      Update a Group
// @Description  Update a group
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default  body  payload.UpdateGroup  true  "request body"
// @Param        id       path  string               true  "group ID"
// @Security     ApiKeyAuth
// @Success      204
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [put]
func (g *groupsController) putUpdateGroupByID(c echo.Context) error {
	id := c.Param("id")

	payload := new(payload.UpdateGroup)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	err := g.groupService.Update(c.Request().Context(), id, *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

// deleteGroupByID godoc
// @Summary      Delete Group by ID
// @Description  Delete group by ID
// @Tags         groups
// @Produce      json
// @Param        id  path  string  true  "group ID"
// @Security     ApiKeyAuth
// @Success      204
// @Failure      404  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [delete]
func (g *groupsController) deleteGroupByID(c echo.Context) error {
	id := c.Param("id")

	err := g.groupService.Delete(c.Request().Context(), id)

	if err != nil {
		return newErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

// getGenerateQRCode godoc
// @Summary      Generate QR Code
// @Description  Generate QR Code
// @Tags         groups
// @Produce      image/png
// @Param        id  path  string  true  "group ID"
// @Security     ApiKeyAuth
// @Success      200  {file}    binary
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/generate [get]
func (g *groupsController) getGenerateQRCode(c echo.Context) error {
	id := c.Param("id")

	file, err := g.groupService.GenerateQRCode(c.Request().Context(), id)

	if err != nil {
		return newErrorResponse(err)
	}

	return c.Blob(http.StatusOK, "image/png", file)
}

// putUpdateAddress godoc
// @Summary      Update an Address
// @Description  Update an address
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default  body  payload.UpdateAddress  true  "request body"
// @Security     ApiKeyAuth
// @Success      204
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/addresses/{id} [put]
func (g *groupsController) putUpdateAddress(c echo.Context) error {
	id := c.Param("id")

	payload := new(payload.UpdateAddress)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	err := g.addressService.Update(c.Request().Context(), id, *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

// postCreateProperty godoc
// @Summary      Add a Property
// @Description  Add a Property
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default  body  payload.CreateProperty  true  "request body"
// @Param        id       path  string                  true  "group ID"
// @Security     ApiKeyAuth
// @Success      201  {object}  createPropertyResponse
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/properties [post]
func (g *groupsController) postCreateProperty(c echo.Context) error {
	groupID := c.Param("id")

	payload := new(payload.CreateProperty)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	id, err := g.propertyService.Create(c.Request().Context(), groupID, *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	idResponse := map[string]any{"id": id}
	response := model.NewResponse("success", "property successfully created", idResponse)
	return c.JSON(http.StatusCreated, response)
}

// putUpdateProperty godoc
// @Summary      Update a Property
// @Description  Update a Property
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default     body  payload.UpdateProperty  true  "request body"
// @Param        id          path  string                  true  "group ID"
// @Param        propertyID  path  string  true  "property ID"
// @Security     ApiKeyAuth
// @Success      204
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/properties/{propertyID} [put]
func (g *groupsController) putUpdateProperty(c echo.Context) error {
	propertyID := c.Param("propertyID")

	payload := new(payload.UpdateProperty)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	if err := g.propertyService.Update(c.Request().Context(), propertyID, *payload); err != nil {
		return newErrorResponse(err)
	}
	return c.NoContent(http.StatusNoContent)
}

// deleteProperty godoc
// @Summary      Delete a Property
// @Description  Delete a Property
// @Tags         groups
// @Produce      json
// @Param        id          path  string  true  "group ID"
// @Param        propertyID  path  string  true  "property ID"
// @Security     ApiKeyAuth
// @Success      204
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/properties/{propertyID} [delete]
func (g *groupsController) deleteProperty(c echo.Context) error {
	propertyID := c.Param("propertyID")

	if err := g.propertyService.Delete(c.Request().Context(), propertyID); err != nil {
		return newErrorResponse(err)
	}
	return c.NoContent(http.StatusNoContent)
}

// getGeneratePropertyQRCode godoc
// @Summary      Generate Property QR Code
// @Description  Generate Property QR Code
// @Tags         groups
// @Produce      image/png
// @Param        id  path  string  true  "group ID"
// @Param        propertyID  path  string                  true  "property ID"
// @Security     ApiKeyAuth
// @Success      200  {file}    binary
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/properties/{propertyID}/generate [get]
func (g *groupsController) getGeneratePropertyQRCode(c echo.Context) error {
	propertyID := c.Param("propertyID")

	file, err := g.propertyService.GenerateQRCode(c.Request().Context(), propertyID)

	if err != nil {
		return newErrorResponse(err)
	}

	return c.Blob(http.StatusOK, "image/png", file)
}

// createGroupResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type createGroupResponse struct {
	Status  string `json:"status" extensions:"x-order=0"`
	Message string `json:"message" extensions:"x-order=1"`
	Data    idData `json:"data" extensions:"x-order=2"`
}

type idData struct {
	ID string `json:"id"`
}

// groupsResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type groupsResponse struct {
	Status  string     `json:"status" extensions:"x-order=0"`
	Message string     `json:"message" extensions:"x-order=1"`
	Data    groupsData `json:"data" extensions:"x-order=2"`
}

type groupsData struct {
	Groups []response.Group `json:"groups"`
}

// groupResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type groupResponse struct {
	Status  string    `json:"status" extensions:"x-order=0"`
	Message string    `json:"message" extensions:"x-order=1"`
	Data    groupData `json:"data" extensions:"x-order=2"`
}

type groupData struct {
	Group response.Group `json:"group"`
}

// createPropertyResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type createPropertyResponse struct {
	Status  string `json:"status" extensions:"x-order=0"`
	Message string `json:"message" extensions:"x-order=1"`
	Data    idData `json:"data" extensions:"x-order=2"`
}
