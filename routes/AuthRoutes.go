package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/controllers"
)

func RouteAuth(app *echo.Echo) {
	app.POST("/signin", controllers.SignUser)
	app.POST("/login", controllers.LoginUser)
}
