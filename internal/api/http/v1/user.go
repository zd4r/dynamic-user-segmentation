package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type userRoutes struct {
}

func newUserRoutes(handler *echo.Group) {
	r := &userRoutes{}

	handler.GET("/user", r.Get)
}

func (r *userRoutes) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
