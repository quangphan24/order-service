package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-service/payload"
)

func (r *Route) UpdateOrder(c echo.Context) error {
	var req payload.UpdateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := r.useCase.OrderUseCase.UpdateOrder(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Successfully!")
}
