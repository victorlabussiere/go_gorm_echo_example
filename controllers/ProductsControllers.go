package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/services"
)

func AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	product, err := services.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, product)
}

func GetAllProducts(ctx echo.Context) error {
	products, err := services.GetAllProducts(ctx.Request().Context())
	if err != nil {
		log.Fatal(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, products)
}

func GetProductById(ctx echo.Context) error {
	var paramId = ctx.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	result, err := services.GetProductById(ctx.Request().Context(), uint(ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

func GetProductByCategoryId(ctx echo.Context) error {
	var paramCategoryId = ctx.Param("id")
	ID, err := strconv.Atoi(paramCategoryId)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	result, err := services.GetProductByCategoryId(ctx.Request().Context(), uint(ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, result)
}
