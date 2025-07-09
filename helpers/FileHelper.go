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

func DeleteFile(filePath string) error{
	if filePath == "" {
		return fmt.Errorf("file path kosong")
	}

	cleanPath := filepath.Clean(filePath)
	baseDir := filepath.Join("uploads")
	
	relPath, err := filepath.Rel(baseDir, cleanPath)
	if err != nil || relPath == ".." || len(relPath) >= 2 && relPath[0:2] == ".." {
		return fmt.Errorf("invalid file path")
	}

	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found")
	}

	if err := os.Remove(cleanPath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
