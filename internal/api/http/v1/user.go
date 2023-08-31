package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

type userRoutes struct {
	userService       userService
	experimentService experimentService
}

func newUserRoutes(handler *echo.Group, userService userService, experimentService experimentService) {
	r := &userRoutes{
		userService:       userService,
		experimentService: experimentService,
	}

	handler.POST("/user", r.Create)
	handler.DELETE("/user/:id", r.Delete)

	handler.GET("/user/:id/segments", r.GetSegments)
	//handler.POST("/user/:id/segments", r.UpdateUserSegments)
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

	user := &userModel.User{
		Id: userId,
	}

	if err := r.userService.Delete(ctx, user); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

type getUserSegmentsResponse struct {
	Segments []segmentModel.Segment `json:"segments"`
}

func (r *userRoutes) GetSegments(c echo.Context) error {
	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrInvalidUserId.Error())
	}

	user := &userModel.User{
		Id: userId,
	}

	segments, err := r.userService.GetSegments(ctx, user)
	if err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := getUserSegmentsResponse{
		Segments: segments,
	}

	return c.JSON(http.StatusOK, resp)
}
