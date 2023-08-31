package v1

import (
	"log"
	"net/http"
	"strconv"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
)

type userRoutes struct {
	userService       userService
	experimentService experimentService
	segmentService    segmentService
}

func newUserRoutes(handler *echo.Group, userService userService, experimentService experimentService, segmentService segmentService) {
	r := &userRoutes{
		userService:       userService,
		experimentService: experimentService,
		segmentService:    segmentService,
	}

	handler.POST("/user", r.Create)
	handler.DELETE("/user/:id", r.Delete)

	handler.GET("/user/:id/segments", r.GetSegments)
	handler.POST("/user/:id/segments", r.UpdateUserSegments)
}

type createUserRequest struct {
	Id int `json:"id"`
}

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

	experimentsToAdd := make([]*experimentModel.Experiment, len(req.SegmentsToAdd))
	for idx, slug := range req.SegmentsToAdd {
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			log.Println(err) // TODO: logger
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		experiment := &experimentModel.Experiment{
			UserId:    userId,
			SegmentId: segmentWithId.Id,
		}
		experimentsToAdd[idx] = experiment
	}

	if err := r.experimentService.CreateBatch(ctx, experimentsToAdd); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	experimentsToDelete := make([]*experimentModel.Experiment, len(req.SegmentsToRemove))
	for idx, slug := range req.SegmentsToRemove {
		segmentWithId, err := r.segmentService.GetBySlug(ctx, slug)
		if err != nil {
			log.Println(err) // TODO: logger
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		experiment := &experimentModel.Experiment{
			UserId:    userId,
			SegmentId: segmentWithId.Id,
		}
		experimentsToDelete[idx] = experiment
	}

	if err := r.experimentService.DeleteBatch(ctx, experimentsToDelete); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
