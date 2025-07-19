package services

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/mapper"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type TagService interface{
	Create(request models.Tags)(dto.TagResponseDTO, error)
	GetAll()([]dto.TagResponseDTO, error)
	GetDetail(id uint64)(dto.TagDetailResponseDTO, error)
	Update(id uint64, request models.Tags)(dto.TagResponseDTO, error)
	Delete(id uint64)(dto.TagResponseDTO, error)
}

type tagService struct {
	tagRepo repository.TagRepository
	viewRepo repository.ViewBlogRepository
}

func NewTagService(tagRepo repository.TagRepository, viewRepo repository.ViewBlogRepository)  TagService{
	return &tagService{tagRepo: tagRepo, viewRepo: viewRepo}
}

func (s *tagService) Create(request models.Tags)(dto.TagResponseDTO, error){
	tag, err := s.tagRepo.Create(request)
	if err != nil {
		return dto.TagResponseDTO{}, err
	}
	response := mapper.TagRespose(tag)
	return response, nil
}

func (s *tagService) GetAll()([]dto.TagResponseDTO,error){
	tags, err := s.tagRepo.GetAll()
	if err != nil{
		return []dto.TagResponseDTO{},err
	}
	var response []dto.TagResponseDTO
	for _,tag := range tags {
		response = append(response, mapper.TagRespose(tag))
	}

	return response, nil
}

func (s *tagService) GetDetail(id uint64)(dto.TagDetailResponseDTO,error){
	tag, err := s.tagRepo.GetById(id)
	if err != nil{
		return dto.TagDetailResponseDTO{},err
	}
	var blogs []dto.BlogResponseDTO
	for _,blog := range tag.Blogs {
		 viewCount,_ := s.viewRepo.GetCountView(int64(blog.ID))
		blogs = append(blogs, mapper.BlogResponse(blog, int64(viewCount)))
	}
	response := mapper.TagDetailResponse(tag,blogs)
	
	return response, nil
}


func (s *tagService) Update(id uint64, request models.Tags)(dto.TagResponseDTO,error){
	tag, err := s.tagRepo.Update(uint64(id), request)
	if err != nil {
		return dto.TagResponseDTO{}, err
	}
	response := mapper.TagRespose(tag)
	return response,nil
}

func (s *tagService) Delete(id uint64)(dto.TagResponseDTO, error){
		tag, err := s.tagRepo.Delete(uint64(id))
	if err != nil {
		return dto.TagResponseDTO{}, err
	}
	response := mapper.TagRespose(tag)
	return response,nil
}