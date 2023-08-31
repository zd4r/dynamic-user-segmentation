package v1

import (
	_ "github.com/zd4r/dynamic-user-segmentation/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// NewRouter creates new router
// Swagger spec:
// @title       Dynamic user segmentation API
// @description Dynamic user segmentation service. Stores users and segments they belong to.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(userService userService, segmentService segmentService, experimentService experimentService, reportService reportService) *echo.Echo {
	handler := echo.New()

	// Middleware
	handler.Use(middleware.Recover())

	// Swagger
	handler.GET("/docs/*", echoSwagger.WrapHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newUserRoutes(h, userService, experimentService, segmentService, reportService)
		newSegmentRoutes(h, segmentService)
	}

	return handler
}
