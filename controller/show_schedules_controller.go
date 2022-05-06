package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/middleware"
	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/showschedule"
	"github.com/labstack/echo/v4"
)

type showSchedulesController struct {
	service showschedule.ShowScheduleService
}

func NewShowSchedulesController(service showschedule.ShowScheduleService) *showSchedulesController {
	return &showSchedulesController{service: service}
}

func (s *showSchedulesController) Route(e *echo.Group) {
	group := e.Group("/shows", middleware.JWTMiddleware())
	group.POST("", s.postCreateShowSchedule)
}

// postCreateShowSchedule godoc
// @Summary      Create a Show Schedule
// @Description  Create a new show schedule
// @Tags         shows
// @Accept       json
// @Produce      json
// @Param        default  body  payload.CreateShowSchedule  true  "request body"
// @Security     ApiKeyAuth
// @Success      201  {object}  createShowScheduleResponse
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /shows [post]
func (s *showSchedulesController) postCreateShowSchedule(c echo.Context) error {
	payload := new(payload.CreateShowSchedule)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	id, err := s.service.Create(c.Request().Context(), *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	idResponse := map[string]any{"id": id}
	response := model.NewResponse("success", "show schedule successfully created", idResponse)
	return c.JSON(http.StatusCreated, response)
}

// createGroupResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type createShowScheduleResponse struct {
	Status  string `json:"status" extensions:"x-order=0"`
	Message string `json:"message" extensions:"x-order=1"`
	Data    idData `json:"data" extensions:"x-order=2"`
}
