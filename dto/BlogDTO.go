package dto

type BlogDTO struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Image   string `json:"image"`
	UserID  uint64 `json:"user_id" form:"user_id" binding:"required"`
}

type BlogResponseDTO struct {
	ID      uint64 ` json:"id"`
	Title   string ` json:"title" `
	Content string ` json:"content" `
	Image   string ` json:"image"`
	Author  string `json:"author"`
	Tags    []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
}

type BlogDetailResponseDTO struct {
	ID      uint64 ` json:"id"`
	Title   string ` json:"title" `
	Content string ` json:"content" `
	Image   string ` json:"image"`
	Author  string `json:"author"`
	Tags    []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	Comment []CommentResponseDTO `json:"comments"`
}