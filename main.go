package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saaitt/user_authentication/handler/user"
)

func main() {


	e := echo.New()
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	userHandler := user.MakeUserHandler()
	user.Api(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))

}