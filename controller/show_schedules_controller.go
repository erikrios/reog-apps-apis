package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/middleware"
	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
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
	group.GET("", s.getShowSchedules)
	group.GET("/:id", s.getShowScheduleByID)
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

// getShowSchedules    godoc
// @Summary      Get Show Schedules
// @Description  Get show schedules
// @Tags         shows
// @Produce      json
// @Param        group_id  query  string  false  "filter show schedules by group ID"
// @Security     ApiKeyAuth
// @Success      200  {object}  showSchedulesResponse
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /shows [get]
func (s *showSchedulesController) getShowSchedules(c echo.Context) error {
	groupID := c.QueryParam("group_id")
	var showSchedules []response.ShowSchedule
	var err error

	if groupID == "" {
		showSchedules, err = s.service.GetAll(c.Request().Context())
	} else {
		showSchedules, err = s.service.GetByGroupID(c.Request().Context(), groupID)
	}

	if err != nil {
		return newErrorResponse(err)
	}

	showSchedulesResposes := map[string]any{"shows": showSchedules}
	responses := model.NewResponse("success", "successfully get show schedules", showSchedulesResposes)
	return c.JSON(http.StatusOK, responses)
}

// getShowScheduleByID godoc
// @Summary      Get Show Schedule by ID
// @Description  Get Show Schedule by ID
// @Tags         shows
// @Produce      json
// @Param        id  path  string  true  "show schedule ID"
// @Security     ApiKeyAuth
// @Success      200  {object}  showScheduleResponse
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /groups/{id} [get]
func (s *showSchedulesController) getShowScheduleByID(c echo.Context) error {
	id := c.Param("id")

	show, err := s.service.GetByID(c.Request().Context(), id)

	if err != nil {
		return newErrorResponse(err)
	}

	showScheduleResponse := map[string]any{"show": show}
	response := model.NewResponse("success", "successfully get show schedule with id "+id, showScheduleResponse)
	return c.JSON(http.StatusOK, response)
}

// createGroupResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type createShowScheduleResponse struct {
	Status  string `json:"status" extensions:"x-order=0"`
	Message string `json:"message" extensions:"x-order=1"`
	Data    idData `json:"data" extensions:"x-order=2"`
}

// showSchedulesResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type showSchedulesResponse struct {
	Status  string            `json:"status" extensions:"x-order=0"`
	Message string            `json:"message" extensions:"x-order=1"`
	Data    showSchedulesData `json:"data" extensions:"x-order=2"`
}

type showSchedulesData struct {
	ShowSchedules []response.ShowSchedule `json:"shows"`
}

// showScheduleResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type showScheduleResponse struct {
	Status  string           `json:"status" extensions:"x-order=0"`
	Message string           `json:"message" extensions:"x-order=1"`
	Data    showScheduleData `json:"data" extensions:"x-order=2"`
}

type showScheduleData struct {
	ShowSchedule response.ShowScheduleDetails `json:"show"`
}
