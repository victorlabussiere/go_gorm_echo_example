package services

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
)

func UploadFile(c context.Context, file *multipart.FileHeader, content []byte) (*models.File, error) {

	extension := strings.ToLower(filepath.Ext(file.Filename))

	uploadedFile := models.File{
		Name:      file.Filename,
		Extension: extension,
		Size:      file.Size,
		Content:   content,
	}

	if err := initializer.DB.Create(&uploadedFile).Error; err != nil {
		return nil, err
	}

	return &uploadedFile, nil

}

func GetFileById(c context.Context, id string) (*models.File, error) {
	var file models.File
	if err := initializer.DB.WithContext(c).First(&file, id).Error; err != nil {
		return nil, err
	}

	return &file, nil
}
