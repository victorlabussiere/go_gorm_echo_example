package initializer

import (
	"fmt"
	"os"

	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_URL"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database'")
	}
}

func SyncDb() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
}
