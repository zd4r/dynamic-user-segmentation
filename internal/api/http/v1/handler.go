package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(userService userService, segmentService segmentService) *echo.Echo {
	handler := echo.New()
	handler.Use(middleware.Recover())

	h := handler.Group("/v1")
	{
		newUserRoutes(h, userService)
		newSegmentRoutes(h, segmentService)
	}

	return handler
}
