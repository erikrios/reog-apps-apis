package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/model"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/admin"
	"github.com/labstack/echo/v4"
)

type adminsController struct {
	service admin.AdminService
}

func NewAdminsController(service admin.AdminService) *adminsController {
	return &adminsController{service: service}
}

func (a *adminsController) Route(g *echo.Group) {
	group := g.Group("/admins")
	group.POST("", a.postLogin)
}

// PostLogin     godoc
// @Summary      Admin Login
// @Description  Admin login
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        default  body      payload.Credential  true  "admin credentials"
// @Success      200      {object}  loginResponse
// @Failure      400      {object}  echo.HTTPError
// @Failure      401      {object}  echo.HTTPError
// @Failure      404      {object}  echo.HTTPError
// @Failure      500      {object}  echo.HTTPError
// @Router       /admins [post]
func (a *adminsController) postLogin(c echo.Context) error {
	credential := new(payload.Credential)
	if err := c.Bind(credential); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	token, err := a.service.Login(c.Request().Context(), *credential)
	if err != nil {
		return newErrorResponse(err)
	}

	tokenResponse := map[string]any{"token": token}
	response := model.NewResponse("success", "login successful", tokenResponse)
	return c.JSON(http.StatusOK, response)
}

// loginResponse struct is used for swaggo to generate the API documentation, as it doesn't support generic yet.
type loginResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    tokenData `json:"data"`
}

type tokenData struct {
	Token string `json:"token"`
}
