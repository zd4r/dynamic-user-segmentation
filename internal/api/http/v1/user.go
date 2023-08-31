package v1

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
)

type userRoutes struct {
	userService       userService
	experimentService experimentService
	segmentService    segmentService
	reportService     reportService
}

func newUserRoutes(handler *echo.Group, userService userService, experimentService experimentService, segmentService segmentService, reportService reportService) {
	r := &userRoutes{
		userService:       userService,
		experimentService: experimentService,
		segmentService:    segmentService,
		reportService:     reportService,
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
	ctx := c.Request().Context()

	req := new(createUserRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &userModel.User{
		Id: req.Id,
	}

	if err := r.userService.Create(ctx, user); err != nil {
		log.Println(err) // TODO: logger
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
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	if err := r.userService.Delete(ctx, userId); err != nil {
		log.Println(err) // TODO: logger
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
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	segments, err := r.userService.GetSegments(ctx, userId)
	if err != nil {
		log.Println(err) // TODO: logger
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
	SegmentsToAdd    []string `json:"segmentsToAdd"`
	SegmentsToRemove []string `json:"segmentsToRemove"`
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
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	req := new(updateUserSegmentsRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	experimentsToAdd := make([]experimentModel.Experiment, 0, len(req.SegmentsToAdd))
	addRecords := make([]reportModel.Record, 0, len(req.SegmentsToAdd))
	for _, slug := range req.SegmentsToAdd {
		slug := strings.ToLower(slug)
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			log.Println(err) // TODO: logger
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		experiment := experimentModel.Experiment{
			UserId:    userId,
			SegmentId: segmentWithId.Id,
		}
		experimentsToAdd = append(experimentsToAdd, experiment)

		record := reportModel.Record{
			UserId:      userId,
			SegmentSlug: slug,
			Action:      reportModel.AddAction,
		}
		addRecords = append(addRecords, record)
	}

	experimentsToDelete := make([]experimentModel.Experiment, 0, len(req.SegmentsToRemove))
	removeRecords := make([]reportModel.Record, 0, len(req.SegmentsToRemove))
	for _, slug := range req.SegmentsToRemove {
		slug := strings.ToLower(slug)
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			log.Println(err) // TODO: logger
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
		}
		removeRecords = append(removeRecords, record)
	}

	if err := r.experimentService.CreateBatch(ctx, experimentsToAdd); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := r.reportService.CreateBatchRecord(ctx, addRecords); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.experimentService.DeleteBatch(ctx, experimentsToDelete); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := r.reportService.CreateBatchRecord(ctx, removeRecords); err != nil {
		log.Println(err) // TODO: logger
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
// @Param       from    query     string  false  "report from time (inclusive)"
// @Param       to      query     string  false  "report to time (not inclusive)"
// @Success		200
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /user/{id}/report [get]
func (r *userRoutes) GetReport(c echo.Context) error {
	const (
		dateLayout = "2006-01"
	)

	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	from := c.QueryParam("from")
	to := c.QueryParam("to")

	var fromDate, toDate = time.Time{}, time.Now()

	if from != "" {
		fromDate, err = time.Parse(dateLayout, from)
		if err != nil {
			log.Println(err) // TODO: logger
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	if to != "" {
		toDate, err = time.Parse(dateLayout, to)
		if err != nil {
			log.Println(err) // TODO: logger
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	report, err := r.reportService.GetRecordsInIntervalByUser(ctx, userId, fromDate, toDate)
	if err != nil {
		log.Println(err) // TODO: logger
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
