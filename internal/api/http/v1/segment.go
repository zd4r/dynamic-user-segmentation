package v1

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	"go.uber.org/zap"
)

type segmentRoutes struct {
	segmentService    segmentService
	experimentService experimentService
	userService       userService
	reportService     reportService
	log               *zap.Logger
}

func newSegmentRoutes(handler *echo.Group, segmentService segmentService, experimentService experimentService, userService userService, reportService reportService, log *zap.Logger) {
	r := &segmentRoutes{
		segmentService:    segmentService,
		experimentService: experimentService,
		userService:       userService,
		reportService:     reportService,
		log:               log.Named("SegmentRoutes"),
	}

	handler.POST("/segment", r.Create)
	handler.DELETE("/segment/:slug", r.Delete)
}

type createSegmentRequest struct {
	Slug         string `json:"slug"`
	UsersPercent *int   `json:"usersPercent"`
}

// Create creates new segment
// @Summary     Create new segment
// @Description Create new segment for users to be put in
// @Tags        segment
// @ID          create-segment
// @Accept      json
// @Param       createSegmentRequest body createSegmentRequest true "Provide slug of a new segment, optionally provide usersPercent [0,100] (%)"
// @Success     201
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /segment [post]
func (r *segmentRoutes) Create(c echo.Context) error {
	l := r.log.Named("Create")
	ctx := c.Request().Context()
	var err error

	req := new(createSegmentRequest)
	if err := c.Bind(req); err != nil {
		l.Error("c.Bind", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var users []userModel.User
	if req.UsersPercent != nil {
		if *req.UsersPercent < 0 || *req.UsersPercent > 100 {
			l.Error("req.UsersPercent", zap.Error(errs.ErrInvalidUserPercent))
			return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserPercent)
		}

		users, err = r.userService.GetPercentOfAllUsers(ctx, *req.UsersPercent)
		if err != nil {
			l.Error("r.userService.GetPercentOfAllUsers", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	slug := strings.ToLower(req.Slug)

	if err := r.segmentService.Create(ctx, slug); err != nil {
		l.Error("r.segmentService.Create", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if req.UsersPercent != nil {
		segment, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			l.Error("r.segmentService.GetBySlug", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		addRecords := make([]reportModel.Record, 0, len(users))
		experimentsToAdd := make([]experimentModel.Experiment, 0, len(users))
		for _, user := range users {
			experiment := experimentModel.Experiment{
				UserId:    user.Id,
				SegmentId: segment.Id,
			}
			experimentsToAdd = append(experimentsToAdd, experiment)

			record := reportModel.Record{
				UserId:      user.Id,
				SegmentSlug: segment.Slug,
				Action:      reportModel.AddAction,
				Date:        time.Now().UTC(),
			}
			addRecords = append(addRecords, record)
		}

		if err := r.experimentService.CreateBatch(ctx, experimentsToAdd); err != nil {
			l.Error("r.experimentService.CreateBatch", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := r.reportService.CreateBatchRecord(ctx, addRecords); err != nil {
			l.Error("r.reportService.CreateBatchRecord", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
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

	users, err := r.segmentService.GetUsersBySlug(ctx, slug)
	if err != nil {
		l.Error("r.segmentService.GetUsersBySlug", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	removeRecords := make([]reportModel.Record, 0, len(users))
	for _, user := range users {
		record := reportModel.Record{
			UserId:      user.Id,
			SegmentSlug: slug,
			Action:      reportModel.RemoveAction,
			Date:        time.Now().UTC(),
		}
		removeRecords = append(removeRecords, record)
	}

	if err := r.segmentService.DeleteBySlug(ctx, slug); err != nil {
		l.Error("r.segmentService.DeleteBySlug", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.reportService.CreateBatchRecord(ctx, removeRecords); err != nil {
		l.Error("r.reportService.CreateBatchRecord", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
