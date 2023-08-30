package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

type userRoutes struct {
	userService userService
}

func newUserRoutes(handler *echo.Group, userService userService) {
	r := &userRoutes{
		userService: userService,
	}

	handler.POST("/user", r.Create)
	handler.DELETE("/user/:id", r.Delete)
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
		log.Println(err)
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
