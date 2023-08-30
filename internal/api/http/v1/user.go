package v1

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	handler.DELETE("/user", r.Delete)
}

type createUserRequest struct {
	Login string `json:"login"`
}

func (r *userRoutes) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(createUserRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &userModel.User{
		Login: req.Login,
	}

	if err := r.userService.Create(ctx, user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

type deleteUserRequest struct {
	Login string `json:"login"`
}

func (r *userRoutes) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(deleteUserRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &userModel.User{
		Login: req.Login,
	}

	if err := r.userService.Delete(ctx, user); err != nil {
		log.Println(err) // TODO: logger
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
