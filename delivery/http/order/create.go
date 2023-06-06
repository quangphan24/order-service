package order

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"order-service/payload"
)

func (r *Route) Create(c echo.Context) error {
	var req payload.CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser, err := r.useCase.OrderUseCase.CreateOrder(c.Request().Context(), req)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}
