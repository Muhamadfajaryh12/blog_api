package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type ViewBlogRepository interface {
	Create(viewBlog models.ViewBlog) (models.ViewBlog, error)
	GetCountView(id int64)(int64, error)
}

type viewBlogRepo struct {
	db *gorm.DB
}

func NewViewBlogRepository(db *gorm.DB) ViewBlogRepository{
	return &viewBlogRepo{db:db}
}

func (r* viewBlogRepo) Create(viewBlog models.ViewBlog)(models.ViewBlog, error){
	err := r.db.Create(&viewBlog).Error
	if err != nil{
		return viewBlog, err
	}
	return viewBlog,nil
}


func (r* viewBlogRepo) GetCountView(id int64)(int64, error){
	var count int64
	err := r.db.Model(&models.ViewBlog{}).Where("blog_id = ?",id).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
