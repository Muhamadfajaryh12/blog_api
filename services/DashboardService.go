package services

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type DashboardService interface {
	GetDashboard(id uint64)(dto.DashboardDTO, error)
}

type dashboardService struct {
	dashboardRepo repository.DashboardRepository
}

func NewDashboardService(dashboardRepo repository.DashboardRepository) DashboardService{
	return &dashboardService{dashboardRepo: dashboardRepo}
}

func (s *dashboardService) GetDashboard(id uint64)(dto.DashboardDTO, error){
	count_view,err := s.dashboardRepo.CountViewAll(id)
	if err != nil{
		return dto.DashboardDTO{}, err
	}

	count_blog,err := s.dashboardRepo.CountBlog(id)
		if err != nil{
		return dto.DashboardDTO{}, err
	}

	count_comment,err := s.dashboardRepo.CountComment(id)
		if err != nil{
		return dto.DashboardDTO{}, err
	}

	count_view_week, err := s.dashboardRepo.CountViewWeek(id)
	if err != nil{
		return dto.DashboardDTO{}, err
	}

	
	response := dto.DashboardDTO{
		CountView: count_view,
		CountBlog: count_blog,
		CountComment: count_comment,
		CoutViewWeek: count_view_week,
	}

	return response,nil
	
}