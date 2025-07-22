package services

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/mapper"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type BlogService interface{
	Create(request models.Blogs)(dto.BlogDetailResponseDTO, error)
	GetAll()(map[string][]dto.BlogResponseDTO, error)
	GetDetail(id uint64)(dto.BlogDetailResponseDTO, error)
	Update(id uint64, request models.Blogs)(dto.BlogDetailResponseDTO, error)
	Delete(id uint64)(dto.BlogDetailResponseDTO, error)
	Search(keyword string)([]dto.BlogResponseDTO, error)
}

type blogService struct {
	blogRepo repository.BlogRepository
	viewRepo repository.ViewBlogRepository
}

func NewBlogService(blogRepo repository.BlogRepository, viewRepo repository.ViewBlogRepository) BlogService {
	return &blogService{blogRepo: blogRepo, viewRepo: viewRepo}
}

func (s *blogService) GetAll() (map[string][]dto.BlogResponseDTO, error){
	blogsAll, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, err
	}

	blogsTrending, err := s.blogRepo.GetTrending()
	if err != nil {
		return nil, err
	}
	blogsLatest, err := s.blogRepo.GetLatest()
	if err != nil {
		return nil, err
	}

	mapBlogs := func(blogs []models.Blogs) []dto.BlogResponseDTO {
		var result []dto.BlogResponseDTO
		for _, blog := range blogs {
			viewCount, _ := s.viewRepo.GetCountView(int64(blog.ID))
			result = append(result, mapper.BlogResponse(blog, viewCount))
		}
		return result
	}

	return map[string][]dto.BlogResponseDTO{
		"all":      mapBlogs(blogsAll),
		"trending": mapBlogs(blogsTrending),
		"latest":   mapBlogs(blogsLatest),
	}, nil
}

func (s *blogService) GetDetail(id uint64) (dto.BlogDetailResponseDTO, error){
	
	blog, err := s.blogRepo.GetDetail(uint64(id))
	if err != nil {
		return dto.BlogDetailResponseDTO{}, err
	}

	viewBlog := models.ViewBlog{
		BlogID: blog.ID,
		View: 1,
	}
	_, err = s.viewRepo.Create(viewBlog)
	if err != nil {
		return dto.BlogDetailResponseDTO{},  err
	}

	viewCount, err := s.viewRepo.GetCountView(int64(blog.ID))
	if err != nil {
		return  dto.BlogDetailResponseDTO{},  err
	}

	result := mapper.BlogDetailResponse(blog,viewCount)
	return result, nil
}

func (s *blogService) Create(request models.Blogs) (dto.BlogDetailResponseDTO, error){
	blog, err := s.blogRepo.Create(request)
	if err != nil {
		return dto.BlogDetailResponseDTO{}, err	
	}
	response := mapper.BlogDetailResponse(blog, 0)
	return response,nil
}

func (s *blogService) Update(id uint64, request models.Blogs) (dto.BlogDetailResponseDTO, error){
	blog, err := s.blogRepo.Update(id, request)
	if err != nil {
		return dto.BlogDetailResponseDTO{}, err	
	}

	viewCount, err := s.viewRepo.GetCountView(int64(blog.ID))
	if err != nil {
		return  dto.BlogDetailResponseDTO{},  err
	}

	response := mapper.BlogDetailResponse(blog, viewCount)
	return response,nil
}

func (s *blogService) Delete(id uint64) (dto.BlogDetailResponseDTO, error){
	blog, err := s.blogRepo.Delete(id)
	if err != nil {
		return dto.BlogDetailResponseDTO{},err
	}
	response := mapper.BlogDetailResponse(blog,0)
	return response,nil
}

func (s *blogService) Search(keyword string)([]dto.BlogResponseDTO,error){
	var response []dto.BlogResponseDTO
	blogs,err := s.blogRepo.Search(keyword)
	if err != nil {
		return []dto.BlogResponseDTO{},err
	}

	for _, blog := range blogs {
		viewCount, _ := s.viewRepo.GetCountView(int64(blog.ID))
		response = append(response,mapper.BlogResponse(blog,viewCount))
	}
	return response,nil
}