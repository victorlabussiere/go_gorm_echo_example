package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUser(c echo.Context) error {

	log.Info("Chamada SignIn iniciada")
	if c.Bind(&body) != nil {
		log.Warn("Chamada SignIn falhou ao ler o body da request")
		c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to read body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Warn("Chamada SignIn falhou ao ecnriptografar a senha no banco")
		c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to hash password",
		})
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializer.DB.Create(&user)
	if result.Error != nil {
		log.Warn("Chamada SignIn falhou ao salvar o usuário no banco")
		c.JSON(http.StatusBadRequest, echo.Map{
			"data":  "Failed to create new user",
			"error": true,
		})
	}

	log.Info("Chamada SignIn concluída com sucesso")
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Usuário criado com sucesso",
		"error":   false,
	})

}

func LoginUser(c echo.Context) error {
	log.Info("Chamada LoginUser iniciada")
	var err error

	if c.Bind(&body) != nil {
		log.Warn("Chamada SignIn falhou ao ler o body da request")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to read body",
		})
	}
	var user = models.User{}
	result := initializer.DB.Where(&models.User{Email: body.Email}).First(&user)
	if result.Error != nil {
		log.Warn("Chamada Login falhou ao buscar o usuário pelo e-mail")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"data":  "Usuário e/ou senha estão incorretos. Verifique e tente novamente",
			"error": true,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		log.Warn("Chamada Login falhou ao comparar as senhas")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"data":  "Usuário e/ou senha estão incorretos. Verifique e tente novamente",
			"error": true,
		})
	}

	type MyCustomClaims struct {
		Sub string `json:"sub"`
		jwt.StandardClaims
	}

	claims := MyCustomClaims{
		"sub",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	claims.Sub = user.Email
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		println(err.Error())
		log.Warn("Chamada Login falhou ao criar assinatura do token")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"data":  "Usuário e/ou senha estão incorretos. Verifique e tente novamente",
			"error": true,
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = expirationTime
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"

	c.SetCookie(cookie)
	log.Info("Chamada LoginUser finalizada")
	return c.JSON(http.StatusOK, echo.Map{
		"data":  "usuário autenticado",
		"error": false,
	})
}
