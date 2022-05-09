package user

import (
	"github.com/labstack/echo"
	"github.com/saaitt/user_authentication/repo"
	"gorm.io/gorm"
)

func MakeUserHandler(db *gorm.DB) Handler {
	return Handler{
		Service: Service{
			Repo: &repo.SQLUserRepo{
				DB: db,
			},
		},
	}

}

func Api(e *echo.Echo, userHandler Handler) {
	userGroup := e.Group("/user")

	userGroup.POST("/", userHandler.Create)
	userGroup.GET("/", userHandler.List)
	userGroup.DELETE("/:user_id", userHandler.DELETE)
	userGroup.POST("/signup", userHandler.Signup)
	userGroup.POST("/signup/cancel/:user_id", userHandler.CancelSignup)
}
