package v1

import (
	"database/sql"
	"encoding/csv"
	"net/http"
	"strconv"
	"strings"
	"time"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
)

type userRoutes struct {
	userService       userService
	experimentService experimentService
	segmentService    segmentService
	reportService     reportService
	log               *zap.Logger
}

func newUserRoutes(handler *echo.Group, userService userService, experimentService experimentService, segmentService segmentService, reportService reportService, log *zap.Logger) {
	r := &userRoutes{
		userService:       userService,
		experimentService: experimentService,
		segmentService:    segmentService,
		reportService:     reportService,
		log:               log.Named("UserRoutes"),
	}

	handler.POST("/user", r.Create)
	handler.DELETE("/user/:id", r.Delete)

	handler.GET("/user/:id/segments", r.GetSegments)
	handler.POST("/user/:id/segments", r.UpdateUserSegments)

	handler.GET("/user/:id/report", r.GetReport)
}

type createUserRequest struct {
	Id int `json:"id"`
}

// Create creates new user
// @Summary     Create new user
// @Description Create new user
// @Tags        user
// @ID          create-user
// @Accept      json
// @Param       createUserRequest body createUserRequest true "Contain user id"
// @Success     201
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user [post]
func (r *userRoutes) Create(c echo.Context) error {
	l := r.log.Named("Create")
	ctx := c.Request().Context()

	req := new(createUserRequest)
	if err := c.Bind(req); err != nil {
		l.Error("c.Bind", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &userModel.User{
		Id: req.Id,
	}

	if err := r.userService.Create(ctx, user); err != nil {
		l.Error("r.userService.Create", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// Delete deletes user
// @Summary     Delete user
// @Description Delete user
// @Tags        user
// @ID          delete-user
// @Param       id path int true "User id"
// @Success     200
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user/{id} [delete]
func (r *userRoutes) Delete(c echo.Context) error {
	l := r.log.Named("Delete")
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.Error("strconv.Atoi", zap.Error(errs.ErrInvalidUserId))
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	if err := r.userService.Delete(ctx, userId); err != nil {
		l.Error("r.userService.Delete", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

type getUserSegmentsResponse struct {
	Segments []string `json:"segments"`
}

// GetSegments returns segments in which user consists
// @Summary     Get user's segments
// @Description Get segments in which user consists
// @Tags        user
// @ID          user-segments
// @Produce      json
// @Param       id path int true "User id"
// @Success     200 {object} getUserSegmentsResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user/{id}/segments [get]
func (r *userRoutes) GetSegments(c echo.Context) error {
	l := r.log.Named("GetSegments")
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.Error("strconv.Atoi", zap.Error(errs.ErrInvalidUserId))
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	segments, err := r.userService.GetSegments(ctx, userId)
	if err != nil {
		l.Error("r.userService.GetSegments", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := getUserSegmentsResponse{
		Segments: make([]string, len(segments)),
	}
	for idx, segment := range segments {
		resp.Segments[idx] = segment.Slug
	}

	return c.JSON(http.StatusOK, resp)
}

type updateUserSegmentsRequest struct {
	SegmentsToAdd []struct {
		Slug     string     `json:"slug"`
		ExpireAt *time.Time `json:"expireAt"`
	} `json:"segmentsToAdd"`
	SegmentsToRemove []struct {
		Slug string `json:"slug"`
	} `json:"segmentsToRemove"`
}

// UpdateUserSegments updates segments in which user consists
// @Summary     Update user's segments
// @Description Update segments in which user consists
// @Tags        user
// @ID          update-user-segments
// @Accept      json
// @Param       id path int true "User id"
// @Param       updateUserSegmentsRequest body updateUserSegmentsRequest true "Contain segments to be added and deleted"
// @Success     200
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user/{id}/segments [post]
func (r *userRoutes) UpdateUserSegments(c echo.Context) error {
	l := r.log.Named("UpdateUserSegments")
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.Error("strconv.Atoi", zap.Error(errs.ErrInvalidUserId))
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	req := new(updateUserSegmentsRequest)
	if err := c.Bind(req); err != nil {
		l.Error("c.Bind", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	addRecords := make([]reportModel.Record, 0, len(req.SegmentsToAdd))
	removeRecords := make([]reportModel.Record, 0, len(req.SegmentsToRemove))

	experimentsToAdd := make([]experimentModel.Experiment, 0, len(req.SegmentsToAdd))
	for _, segment := range req.SegmentsToAdd {
		slug := strings.ToLower(segment.Slug)
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			l.Error("r.segmentService.GetBySlug", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		experiment := experimentModel.Experiment{
			UserId:    userId,
			SegmentId: segmentWithId.Id,
		}

		if segment.ExpireAt != nil {
			if !time.Now().UTC().Before(*segment.ExpireAt) {
				l.Error("time.Now().Before", zap.Error(errs.ErrInvalidExpireAtTime))
				return echo.NewHTTPError(http.StatusInternalServerError, errs.ErrInvalidExpireAtTime)
			}
			experiment.ExpireAt = sql.NullTime{
				Time:  *segment.ExpireAt,
				Valid: true,
			}
		}
		experimentsToAdd = append(experimentsToAdd, experiment)

		record := reportModel.Record{
			UserId:      userId,
			SegmentSlug: slug,
			Action:      reportModel.AddAction,
			Date:        time.Now().UTC(),
		}
		addRecords = append(addRecords, record)

		if experiment.ExpireAt.Valid {
			record.Action = reportModel.RemoveAction
			record.Date = experiment.ExpireAt.Time
			removeRecords = append(removeRecords, record)
		}
	}

	experimentsToDelete := make([]experimentModel.Experiment, 0, len(req.SegmentsToRemove))
	for _, segment := range req.SegmentsToRemove {
		slug := strings.ToLower(segment.Slug)
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			l.Error("r.segmentService.GetBySlug", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		experiment := experimentModel.Experiment{
			UserId:    userId,
			SegmentId: segmentWithId.Id,
		}
		experimentsToDelete = append(experimentsToDelete, experiment)

		record := reportModel.Record{
			UserId:      userId,
			SegmentSlug: slug,
			Action:      reportModel.RemoveAction,
			Date:        time.Now().UTC(),
		}
		removeRecords = append(removeRecords, record)
	}

	if err := r.experimentService.CreateBatch(ctx, experimentsToAdd); err != nil {
		l.Error("r.experimentService.CreateBatch", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := r.reportService.CreateBatchRecord(ctx, addRecords); err != nil {
		l.Error("r.reportService.CreateBatchRecord", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.experimentService.DeleteBatch(ctx, experimentsToDelete); err != nil {
		l.Error("r.experimentService.DeleteBatch", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := r.reportService.CreateBatchRecord(ctx, removeRecords); err != nil {
		l.Error("r.reportService.CreateBatchRecord", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

// GetReport creates report in CSV format with user actions within a given time range
// @Summary     Create report in CSV format
// @Description Create report in CSV format with user actions within a given time range
// @Tags        user
// @ID          get-user-report
// @Param       id path int true "User id"
// @Param       from    query     string  false  "report from date [inclusive] (format: YYYY-MM)"
// @Param       to      query     string  false  "report to date [not inclusive] (format: YYYY-MM)"
// @Success		200
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user/{id}/report [get]
func (r *userRoutes) GetReport(c echo.Context) error {
	const (
		dateLayout = "2006-01"
	)

	l := r.log.Named("GetReport")
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		l.Error("strconv.Atoi", zap.Error(errs.ErrInvalidUserId))
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	from := c.QueryParam("from")
	to := c.QueryParam("to")

	var fromDate, toDate = time.Time{}, time.Now().UTC()

	if from != "" {
		fromDate, err = time.Parse(dateLayout, from)
		if err != nil {
			l.Error("time.Parse", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	if to != "" {
		toDate, err = time.Parse(dateLayout, to)
		if err != nil {
			l.Error("time.Parse", zap.Error(err))
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	report, err := r.reportService.GetRecordsInIntervalByUser(ctx, userId, fromDate, toDate)
	if err != nil {
		l.Error("r.reportService.GetRecordsInIntervalByUser", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	reportCSV := make([][]string, 0, len(report)+1)
	reportCSV = append(reportCSV, reportModel.GetCSVHeader())
	for _, record := range report {
		reportCSV = append(reportCSV, record.ToCSV())
	}

	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment;filename=report.csv")

	w := csv.NewWriter(c.Response().Writer)

	return w.WriteAll(reportCSV)
}
