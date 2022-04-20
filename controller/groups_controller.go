package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/group"
	"github.com/labstack/echo/v4"
)

type groupsController struct {
	service group.GroupService
}

func NewGroupsController(service group.GroupService) *groupsController {
	return &groupsController{service: service}
}

func (g *groupsController) Route(e *echo.Group) {
	group := e.Group("/groups")
	group.POST("", g.postCreateGroup)
	group.GET("", g.getGroups)
	group.GET("/:id", g.getGroupByID)
	group.PUT("/:id", g.putUpdateGroupByID)
	group.DELETE("/:id", g.deleteGroupByID)
	group.GET("/:id/generate", g.getGenerateQRCode)
}

// postCreateGroup godoc
// @Summary      Create a Group
// @Description  Create a new group
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        default  body      payload.CreateGroup  true  "request body"
// @Success      201      {object}  createGroupResponse
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

	id, err := g.service.Create(c.Request().Context(), *payload)
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
// @Success      200  {object}  groupsResponse
// @Failure      401      {object}  echo.HTTPError
// @Failure      500      {object}  echo.HTTPError
// @Router       /groups [get]
func (g *groupsController) getGroups(c echo.Context) error {
	groups, err := g.service.GetAll(c.Request().Context())
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
// @Param        id   path      string  true  "group ID"
// @Success      200  {object}  groupResponse
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [get]
func (g *groupsController) getGroupByID(c echo.Context) error {
	id := c.Param("id")

	group, err := g.service.GetByID(c.Request().Context(), id)

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
// @Success      204
// @Failure      400      {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404      {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [put]
func (g *groupsController) putUpdateGroupByID(c echo.Context) error {
	id := c.Param("id")

	payload := new(payload.UpdateGroup)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	err := g.service.Update(c.Request().Context(), id, *payload)
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
// @Param        id   path      string  true  "group ID"
// @Success      204
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [delete]
func (g *groupsController) deleteGroupByID(c echo.Context) error {
	id := c.Param("id")

	err := g.service.Delete(c.Request().Context(), id)

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
// @Success      200  {file}    binary
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id}/generate [get]
func (g *groupsController) getGenerateQRCode(c echo.Context) error {
	id := c.Param("id")

	file, err := g.service.GenerateQRCode(c.Request().Context(), id)

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
