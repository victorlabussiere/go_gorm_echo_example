package helpers

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

func ReadAsBlob(file multipart.File) ([]byte, error) {
	content := make([]byte, 0)
	buffer := make([]byte, 1024) // Buffer de 1 KB

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		content = append(content, buffer[:n]...)
	}
	return content, nil
}

func GetMimeType(extension string) string {
	switch extension {
	case ".pdf":
		return "application/pdf"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".txt":
		return "text/plain"
	case ".html":
		return "text/html"
	default:
		return "application/octet-stream" // Tipo genérico para arquivos desconhecidos
	}
}

func FileResponseAdapter(c echo.Context, filename string, mimeType string) {
	c.Response().Header().Set(echo.HeaderContentType, mimeType)
	c.Response().Header().Set(
		echo.HeaderContentDisposition,
		fmt.Sprintf("inline; filename=%q", filename),
	)
}
