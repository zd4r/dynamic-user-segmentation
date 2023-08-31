package v1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	"go.uber.org/zap"
)

type segmentRoutes struct {
	segmentService segmentService
	log            *zap.Logger
}

func newSegmentRoutes(handler *echo.Group, segmentService segmentService, log *zap.Logger) {
	r := &segmentRoutes{
		segmentService: segmentService,
		log:            log.Named("SegmentRoutes"),
	}

	handler.POST("/segment", r.Create)
	handler.DELETE("/segment/:slug", r.Delete)
}

type createSegmentRequest struct {
	Slug string `json:"slug"`
}

// Create creates new segment
// @Summary     Create new segment
// @Description Create new segment for users to be put in
// @Tags        segment
// @ID          create-segment
// @Accept      json
// @Param       createSegmentRequest body createSegmentRequest true "Contain segment slug"
// @Success     201
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /segment [post]
func (r *segmentRoutes) Create(c echo.Context) error {
	l := r.log.Named("Create")
	ctx := c.Request().Context()

	req := new(createSegmentRequest)
	if err := c.Bind(req); err != nil {
		l.Error("c.Bind", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	segment := &segmentModel.Segment{
		Slug: strings.ToLower(req.Slug),
	}

	if err := r.segmentService.Create(ctx, segment); err != nil {
		l.Error("r.segmentService.Create", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// Delete deletes segment
// @Summary     Delete segment
// @Description Delete segment
// @Tags        segment
// @ID          delete-segment
// @Param       slug path string true "Segment slug"
// @Success     200
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /segment/{slug} [delete]
func (r *segmentRoutes) Delete(c echo.Context) error {
	l := r.log.Named("Delete")
	ctx := c.Request().Context()

	slug := strings.ToLower(c.Param("slug"))

	if err := r.segmentService.DeleteBySlug(ctx, slug); err != nil {
		l.Error("r.segmentService.DeleteBySlug", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
