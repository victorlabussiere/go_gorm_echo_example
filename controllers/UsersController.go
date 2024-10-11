package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
)

var body struct {
	Email    string
	Password string
}

func IndexUser(c echo.Context) error {
	log.Info("Chamada de busca iniciada")
	var users models.User
	result := initializer.DB.Find(&models.User{}).Scan(&users)
	if result.Error != nil {
		log.Warn("Problemas na consulta")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"data":  "Erro durante a busca dos usuários",
			"error": true,
		})
	}

	log.Info("Chamada de busca concluída com sucesso")
	return c.JSON(http.StatusOK, echo.Map{
		"data":  users,
		"error": false,
	})
}
