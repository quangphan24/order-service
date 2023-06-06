package order

import (
	"github.com/labstack/echo/v4"
	"order-service/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.POST("/create", r.Create)
	group.PUT("/:id", r.UpdateOrder)
}
