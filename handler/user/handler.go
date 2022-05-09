package user

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	Service
}

func (h Handler) Create(c echo.Context) error {
	req := CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	user, err := h.Service.CreateUser(req)
	if err != nil {
		if err != nil {
			return echo.ErrInternalServerError
		}
	}
	return c.JSON(http.StatusOK, user)
}

func (h Handler) List(c echo.Context) error {
	cancelledSubStr := c.QueryParam("cancelled_subscription")
	pageStr := c.QueryParam("page")

	users, err := h.Service.ListUsers(pageStr, cancelledSubStr)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, users)
}


func (h Handler) DELETE(c echo.Context) error {
	userID := c.Param("user_id")

	err := h.Service.DeleteUser(userID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.NoContent(http.StatusNoContent)
}

func (h Handler) Signup(c echo.Context) error {
	req := CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	//TODO: Some kind of validation should be here
	user, err := h.Service.CreateUser(req)
	if err != nil {
		if err != nil {
			return echo.ErrInternalServerError
		}
	}
	return c.JSON(http.StatusOK, user)
}