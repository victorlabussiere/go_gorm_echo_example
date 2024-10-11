package routes

import "github.com/labstack/echo/v4"

func Router(app *echo.Echo) {
	RouteAuth(app)
	RouteUsers(app)
	RouteCategories(app)
	RouteProducts(app)
}
