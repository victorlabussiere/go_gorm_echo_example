package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/routes"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	initializer.SyncDb()
}

func main() {

	// setup
	app := echo.New()
	// routes
	routes.Router(app)
	// start server
	app.Start(":" + os.Getenv("PORT"))

}
