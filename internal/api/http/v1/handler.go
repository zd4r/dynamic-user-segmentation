package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	handler := echo.New()
	handler.Use(middleware.Recover())

	h := handler.Group("/v1")
	{
		newUserRoutes(h)
		newSegmentRoutes(h)
	}

	return handler
}
