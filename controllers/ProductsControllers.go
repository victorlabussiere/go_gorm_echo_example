package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/services"
	"gorm.io/gorm"
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

func UpdateProductById(ctx echo.Context) error {

	var updateProductDto struct {
		Id         uint    `json:"id"`
		Name       string  `json:"name"`
		Value      float64 `json:"value"`
		CategoryId uint    `json:"categoryId,omitempty"`
	}

	if err := ctx.Bind(&updateProductDto); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": "O formato inserido é inválido",
			"data":    err.Error(),
		})
	}

	idParam := ctx.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || uint(idUint) != updateProductDto.Id {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": "Conflito entre os IDs encontrados. Verifique se o ID no parâmetro é o mesmo do body.",
		})
	}

	var product models.Product
	result := initializer.DB.WithContext(ctx.Request().Context()).First(&product, idUint)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"error":   true,
				"message": "Produto não encontrado",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": "Erro ao buscar produto",
			"data":    result.Error.Error(),
		})
	}

	if updateProductDto.CategoryId != 0 {
		var category models.Category
		if err := initializer.DB.WithContext(ctx.Request().Context()).First(&category, updateProductDto.CategoryId).Error; err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"error":   true,
				"message": "A categoria inserida não existe",
			})
		}
		product.CategoryId = updateProductDto.CategoryId
	}

	product.Name = updateProductDto.Name
	product.Value = updateProductDto.Value

	if err := initializer.DB.WithContext(ctx.Request().Context()).Save(&product).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": "Erro ao atualizar produto",
			"data":    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"error":   false,
		"message": "Operação concluida.",
		"data":    product,
	})
}

func DeleteProductById(c echo.Context) error {
	id := c.Param("id")
	result := initializer.DB.Where("id = ?", id).Delete(&models.Product{})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": "Não foi possível concluir a operação",
			"data":    result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error":   true,
			"message": "Produto não encontrado",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error":   false,
		"message": "Produto deletado com sucesso",
	})
}
