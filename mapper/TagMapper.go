package mapper

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/models"
)

func TagRespose(tag models.Tags) dto.TagResponseDTO{
	return dto.TagResponseDTO{
			ID:tag.ID,
			Tag:tag.Tag,
		}	
}

func TagDetailResponse(tag models.Tags, blogs []dto.BlogResponseDTO) dto.TagDetailResponseDTO{
	return dto.TagDetailResponseDTO{
		ID:tag.ID,
		Tag:tag.Tag,
		Blogs: blogs,
	}
}