package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type segmentRoutes struct {
}

func newSegmentRoutes(handler *echo.Group) {
	r := &segmentRoutes{}

	handler.GET("/segment", r.Get)
}

func (r *segmentRoutes) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
