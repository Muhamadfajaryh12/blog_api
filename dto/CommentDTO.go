package dto

type CommentRequestDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content" form:"content" binding:"required" example:"content"`
	BlogID  uint64 `json:"blog_id" form:"blog_id" binding:"required" example:"4"`
}

type CommentResponseDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}