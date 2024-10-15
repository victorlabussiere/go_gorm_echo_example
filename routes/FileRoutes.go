package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/controllers"
)

func RouteFiles(app *echo.Echo) {
	app.POST("/files", controllers.UploadFile)
	app.GET("/files", controllers.RetrieveAllFiles)
	app.GET("/files/:id", controllers.RetrieveFile)
	app.GET("/files/content/:id", controllers.RetrieveFileContent)
}
