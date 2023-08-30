package v1

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

type segmentRoutes struct {
	segmentService segmentService
}

func newSegmentRoutes(handler *echo.Group, segmentService segmentService) {
	r := &segmentRoutes{
		segmentService: segmentService,
	}

	handler.POST("/segment", r.Create)
	handler.DELETE("/segment", r.Delete)
}

type createSegmentRequest struct {
	Slug string `json:"slug"`
}

func (r *segmentRoutes) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(createSegmentRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	segment := &segmentModel.Segment{
		Slug: req.Slug,
	}

	if err := r.segmentService.Create(ctx, segment); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

type deleteSegmentRequest struct {
	Slug string `json:"slug"`
}

func (r *segmentRoutes) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(deleteSegmentRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	segment := &segmentModel.Segment{
		Slug: req.Slug,
	}

	if err := r.segmentService.Delete(ctx, segment); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}
