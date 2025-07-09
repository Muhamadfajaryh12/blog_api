package dto

type CommentDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content" form:"content" binding:"required"`
	UserID  uint64 `json:"user_id" form:"user_id" binding:"required"`
	BlogID  uint64 `json:"blod_id" form:"blog_id" binding:"required"`
}

type CommentResponseDTO struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}