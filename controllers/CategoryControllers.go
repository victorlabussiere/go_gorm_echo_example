package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/services"
)

func AddCategory(ctx echo.Context) error {
	var category = new(models.Category)

	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	category, err := services.AddCategory(ctx.Request().Context(), category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, category)
}

func GetCategories(ctx echo.Context) error {
	categories, err := services.GetCategories(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, categories)
}

func GetCategoryById(ctx echo.Context) error {
	paramId := ctx.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	result, err := services.GetCategoriesById(ctx.Request().Context(), uint(ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}

	var response struct {
		Id   uint
		Name string
	}
	response.Id = result.ID
	response.Name = result.Name

	return ctx.JSON(http.StatusOK, response)
}
