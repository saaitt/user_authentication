package user

import "github.com/labstack/echo"

type Handler struct {

}

func (h Handler)CreateUser(c echo.Context) error{
	req := CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	return nil
}