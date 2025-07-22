package dto

import (
	"mime/multipart"
	"time"
)

type BlogRequestDTO struct {
	ID      uint64                `json:"id"`
	Title   string                `json:"title" form:"title" binding:"required" example:"title blog"`
	Content string                `json:"content" form:"content" binding:"required" example:"content blog"`
	Upload   *multipart.FileHeader `form:"upload" json:"-"`
}

type BlogResponseDTO struct {
	ID      uint64 ` json:"id"`
	Title   string ` json:"title" `
	Content string ` json:"content" `
	Image   string ` json:"image"`
	Author  string ` json:"author"`
	View 	int64 	`json:"view"`
	Date time.Time `json:"date"`
	Tags    []TagResponseDTO `json:"tags"`
}

type BlogDetailResponseDTO struct {
	ID      uint64 ` json:"id"`
	Title   string ` json:"title" `
	Content string ` json:"content" `
	Image   string ` json:"image"`
	Author  string `json:"author"`
	View 	int64 	`json:"view"`
	Date time.Time `json:"date"`
	Tags    []TagResponseDTO `json:"tags"`
	Comment []CommentResponseDTO `json:"comments"`
}

