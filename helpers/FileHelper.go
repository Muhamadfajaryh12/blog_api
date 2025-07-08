package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveFile(fileHeader *multipart.FileHeader, folder string)(string, error){
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(fileHeader.Filename))
	uploadPath := filepath.Join("uploads",folder)
	filePath := filepath.Join(uploadPath,fileName)

	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	return filePath, nil
}