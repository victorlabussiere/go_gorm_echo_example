package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/controllers"
)

func RouteUsers(app *echo.Echo) {
	app.GET("/users", controllers.IndexUser)
	app.POST("/users", controllers.SignUser)
	app.POST("/login", controllers.LoginUser)
}
