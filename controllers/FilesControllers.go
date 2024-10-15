package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/helpers"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/services"
)

func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": "Erro ao obter o arquivo",
			"data":    err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": "Erro ao abrir o arquivo",
			"data":    err.Error(),
		})
	}
	defer src.Close()

	content, err := helpers.ReadAsBlob(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao ler o conteúdo do arquivo").SetInternal(err)
	}

	uploadedFile, err := services.UploadFile(c.Request().Context(), file, content)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": "Erro ao salvar dados no banco",
			"data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error":    false,
		"message":  "Upload realizado com sucesso",
		"file_id":  uploadedFile.ID,
		"filename": uploadedFile.Name,
	})
}

func RetrieveFile(c echo.Context) error {
	id := c.Param("id")
	var file, err = services.GetFileById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error":   true,
			"message": "Conteúdo não encontrado",
			"data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error":   false,
		"data":    file,
		"message": "sucesso",
	})
}

func RetrieveFileContent(c echo.Context) error {
	id := c.Param("id")

	var file, err = services.GetFileById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error":   true,
			"message": "Conteúdo não encontrado",
			"data":    err.Error(),
		})
	}

	mimeType := helpers.GetMimeType(file.Extension)

	helpers.FileResponseAdapter(c, file.Name, mimeType)

	return c.Blob(http.StatusOK, mimeType, file.Content)
}
