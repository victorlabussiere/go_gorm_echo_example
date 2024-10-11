package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/controllers"
)

func RouteCategories(app *echo.Echo) {
	app.POST("/category", controllers.AddCategory)
	app.GET("/category", controllers.GetCategories)
	app.GET("/category/:id", controllers.GetCategoryById)
}
