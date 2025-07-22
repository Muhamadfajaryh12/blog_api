package mapper

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/models"
)

func BlogResponse(blog models.Blogs, viewCount int64 ) dto.BlogResponseDTO{
	var tags []dto.TagResponseDTO

	for _, t := range blog.Tags {
		tags = append(tags,dto.TagResponseDTO{
			ID: t.ID,
			Tag: t.Tag,
		})
	}

	return dto.BlogResponseDTO{
		ID:blog.ID, 
		Title: blog.Title,
		Content: blog.Content,
		Image: blog.Image,
		Date:blog.CreatedAt,
		Author: blog.Users.Name,
		Tags:tags,
		View:viewCount,
	}
}

func BlogDetailResponse(blog models.Blogs, viewCount int64) dto.BlogDetailResponseDTO{
	
	// Manipulasi data tag
	var tags []dto.TagResponseDTO

	for _, t := range blog.Tags {
		tags = append(tags,dto.TagResponseDTO{
			ID: t.ID,
			Tag: t.Tag,
		})
	}


	var comments []dto.CommentResponseDTO

	for _,c := range blog.Comments {
		comments = append(comments,CommentMapper(c))
	}
	
	return dto.BlogDetailResponseDTO{
		ID:blog.ID, 
		Title:  blog.Title,
		Content: blog.Content,
		Image: blog.Image,
		Date:blog.CreatedAt,
		Author: blog.Users.Name,
		View: viewCount,
		Tags: tags,
		Comment: comments,
	}
}