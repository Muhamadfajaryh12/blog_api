package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(inputBlog models.Blogs)(models.Blogs, error)
	GetAll(blog []models.Blogs)([]models.Blogs, error)
}

type blogRepo struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository{
	return &blogRepo{db:db}
}

func (r *blogRepo) Create(inputBlog models.Blogs)(models.Blogs, error){
	err := r.db.Create(&inputBlog).Error
	if err != nil{
		return inputBlog, err
	}

	return inputBlog, nil
}


func (r *blogRepo) GetAll(blog []models.Blogs)([]models.Blogs, error){
	err:= r.db.Find(&blog).Error
	if err != nil{
		return blog, err
	}

	return blog, nil
}