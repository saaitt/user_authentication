package user

import (
	"github.com/labstack/echo"
)

func MakeUserHandler() Handler {
	return Handler{

	}
}

func Api(e *echo.Echo, userHandler Handler) {
	userGroup := e.Group("/user")

	userGroup.POST("/", userHandler.CreateUser)
}