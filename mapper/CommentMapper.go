package mapper

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/models"
)

func CommentMapper(comments models.Comments) (dto.CommentResponseDTO) {
	return dto.CommentResponseDTO{
		ID: comments.ID,
		Content: comments.Content,
		Sender:comments.Users.Name,
	}
}