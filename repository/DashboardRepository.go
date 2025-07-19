package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"gorm.io/gorm"
)

type DashboardRepository interface{
	CountViewAll(id uint64)(int64, error)
	CountBlog(id uint64)(int64,error)
	CountComment(id uint64)(int64,error)
	CountViewWeek(id uint64)([]dto.ViewDayDTO, error)
}

type dashboardRepo struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository{
	return &dashboardRepo{db:db}
}

func (r *dashboardRepo) CountViewAll(id uint64)(int64, error){
	var total int64
	query := `SELECT COALESCE(SUM(view_blogs.view), 0)
	FROM view_blogs
	LEFT JOIN blogs  ON blogs.id = view_blogs.id
	WHERE blogs.user_id = ?
	`
	err := r.db.Raw(query,id).Scan(&total).Error
	if err != nil {
	return 0, err
	}

	return total, err
}

func (r *dashboardRepo) CountBlog(id uint64)(int64,error){
	var total int64
	query := `SELECT COALESCE(COUNT(id),0)
	FROM blogs 
	WHERE user_id = ?
	`
	err := r.db.Raw(query,id).Scan(&total).Error
	if err != nil {
	return 0, err
	}

	return total, err
}

func (r *dashboardRepo) CountComment(id uint64)(int64, error){
	var total int64
	query := `SELECT COALESCE(COUNT(comments.id),0)
	FROM comments
	LEFT JOIN blogs ON blogs.id = comments.blog_id
	WHERE blogs.user_id = ?
	`

	err := r.db.Raw(query,id).Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, err
}

func (r *dashboardRepo) CountViewWeek(id uint64)([]dto.ViewDayDTO, error){
	var results []dto.ViewDayDTO
		query := `
	SELECT
		DATE_FORMAT(d.date, '%b %d') AS date,
		COALESCE(SUM(v.view), 0) AS count_view
	FROM (
		SELECT CURDATE() - INTERVAL n DAY AS date
		FROM (
			SELECT 0 AS n UNION SELECT 1 UNION SELECT 2 UNION SELECT 3
			UNION SELECT 4 UNION SELECT 5 UNION SELECT 6
		) AS numbers
	) AS d
	LEFT JOIN view_blogs v ON DATE(v.created_at) = d.date
	LEFT JOIN blogs b ON b.id = v.blog_id AND b.user_id = ?
	GROUP BY d.date
	ORDER BY d.date ASC
	`
	err := r.db.Raw(query, id).Scan(&results).Error
	if err != nil {
		return []dto.ViewDayDTO{}, err
	}
	return results, err
}