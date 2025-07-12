package mapper

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/models"
)

func BlogResponse(blog models.Blogs) dto.BlogResponseDTO{
	var tags []struct{
		Tag string `json:"tag"`
	}

	for _, t := range blog.Tags {
		tags = append(tags, struct{Tag string "json:\"tag\""}{Tag:t.Tag})
	}

	return dto.BlogResponseDTO{
		ID:blog.ID, 
		Title: blog.Title,
		Content: blog.Content,
		Image: blog.Image,
		View:blog.View,
		Date:blog.CreatedAt,
		Author: blog.Users.Name,
		Tags:tags,
	}
}

func BlogDetailResponse(blog models.Blogs) dto.BlogDetailResponseDTO{
	
	// Manipulasi data tag
	var tags []struct{
		Tag string `json:"tag"`
	}

	for _, t := range blog.Tags {
		tags = append(tags, struct{Tag string "json:\"tag\""}{Tag:t.Tag})
	}

	var comments []dto.CommentResponseDTO

	for _,c := range blog.Comments {
		comments = append(comments, dto.CommentResponseDTO{
			ID: c.ID,
			Content: c.Content,
			Sender: c.Users.Name,
		})
	}
	
	return dto.BlogDetailResponseDTO{
		ID:blog.ID, 
		Title:  blog.Title,
		Content: blog.Content,
		Image: blog.Image,
		View:blog.View,
		Date:blog.CreatedAt,
		Author: blog.Users.Name,
		Tags: tags,
		Comment: comments,
	}
}