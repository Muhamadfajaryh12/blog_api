package dto

type CommentRequestDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content" form:"content" binding:"required" example:"content"`
	UserID  uint64 `json:"user_id" form:"user_id" binding:"required" example:"1"`
	BlogID  uint64 `json:"blod_id" form:"blog_id" binding:"required" example:"2"`
}

type CommentResponseDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}