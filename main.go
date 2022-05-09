package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saaitt/user_authentication/handler/user"
	"github.com/saaitt/user_authentication/repo"
)

func main() {
	db, err := repo.InitDb()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	userHandler := user.MakeUserHandler(db)
	user.Api(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))

}

