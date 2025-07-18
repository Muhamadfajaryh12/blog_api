package services

import "github.com/muhamadfajaryh12/api_blogs/repository"

type TagService interface{}

type tagService struct {
	tagRepo repository.TagRepository
}

func newTagService(tagRepo repository.TagRepository)  TagService{
	return &tagService{tagRepo: tagRepo}
}

