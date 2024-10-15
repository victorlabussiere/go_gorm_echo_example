package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/controllers"
)

func RouteProducts(app *echo.Echo) {
	app.POST("/products", controllers.AddProduct)
	app.GET("/products", controllers.GetAllProducts)
	app.GET("/products/:id", controllers.GetProductById)
	app.GET("/products/category/:id", controllers.GetProductByCategoryId)
	app.PATCH("/products/:id", controllers.UpdateProductById)
	app.DELETE("/products/:id", controllers.DeleteProductById)
}
