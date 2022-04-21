package controller

import (
	"net/http"

	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/service/address"
	"github.com/labstack/echo/v4"
)

type addressesController struct {
	service address.AddressService
}

func NewAddressController(service address.AddressService) *addressesController {
	return &addressesController{service: service}
}

func (a *addressesController) Route(g *echo.Group) {
	group := g.Group("/addresses")
	group.PUT("/:id", a.putUpdateAddress)
}

// putUpdateAddress godoc
// @Summary      Update an Address
// @Description  Update an address
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        default  body  payload.UpdateAddress  true  "request body"
// @Success      204
// @Failure      400  {object}  echo.HTTPError
// @Failure      401  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /addresses/{id} [put]
func (a *addressesController) putUpdateAddress(c echo.Context) error {
	id := c.Param("id")

	payload := new(payload.UpdateAddress)
	if err := c.Bind(payload); err != nil {
		return newErrorResponse(service.ErrInvalidPayload)
	}

	err := a.service.Update(c.Request().Context(), id, *payload)
	if err != nil {
		return newErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}
