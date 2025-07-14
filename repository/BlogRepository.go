package repository

import (
	"strings"

	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(inputBlog models.Blogs)(models.Blogs, error)
	GetAll(blog []models.Blogs)([]models.Blogs, error)
	GetTrending(blog []models.Blogs)([]models.Blogs, error)
	GetLatest(blog []models.Blogs)([]models.Blogs, error)
	GetDetail(id uint64, blog models.Blogs)(models.Blogs, error)
	Update(id uint64, blog models.Blogs)(models.Blogs, error)
	Delete(id uint64)(models.Blogs, error)
	Search(keyword string, blog []models.Blogs)([]models.Blogs, error)
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
	
	if err := r.db.Preload("Users").Preload("Tags").First(&inputBlog, inputBlog.ID).Error; err != nil {
		return inputBlog, err
	}
	return inputBlog, nil
}

func (r *blogRepo) GetAll(blog []models.Blogs)([]models.Blogs, error){
	err:= r.db.Preload("Tags").Preload("Users").Omit("Comments").Order("ID DESC").Limit(6).Find(&blog).Error
	if err != nil{
		return blog, err
	}
	return blog, nil
}

func (r *blogRepo) GetTrending(blog []models.Blogs)([]models.Blogs,error){
	err:= r.db.Preload("Tags").Preload("Users").Omit("Comments").Order("View DESC").Limit(3).Find(&blog).Error
		if err != nil{
		return blog, err
	}
	return blog, nil
}

func (r *blogRepo) GetLatest(blog []models.Blogs)([]models.Blogs,error){
	err:= r.db.Preload("Tags").Preload("Users").Omit("Comments").Limit(3).Find(&blog).Error
		if err != nil{
		return blog, err
	}
	return blog, nil
}

func (r *blogRepo) GetDetail(id uint64,blog models.Blogs)(models.Blogs, error){
	err:= r.db.Preload("Tags").Preload("Users").Preload("Comments.Users").First(&blog,id).Error
	r.db.Model(&blog).UpdateColumn("view", blog.View + 1)
	
	if err != nil{
		return blog, err
	}
	return blog, nil
}

func (r *blogRepo) Update(id uint64, updateData models.Blogs)(models.Blogs, error ){
	var blog models.Blogs
	err:= r.db.Preload("Users").Preload("Tags").First(&blog,id).Error
	if err != nil {
		return blog, err
	}

	blog.Title = updateData.Title
	blog.Content = updateData.Content
	if updateData.Image != "" {
		helpers.DeleteFile(strings.ReplaceAll(blog.Image, "/", "\\"))
		blog.Image = updateData.Image
	}

	err = r.db.Save(&blog).Error;
	if err != nil{
		return blog,err
	}

	err = r.db.Model(&blog).Association("Tags").Replace(updateData.Tags)
	if err != nil{
		return blog,err
	}	

	err= r.db.Preload("Users").Preload("Tags").First(&blog,id).Error
	if err != nil {
		return blog, err
	}
	
	return blog,nil
}

func (r *blogRepo) Delete(id uint64)(models.Blogs, error){
	var blog models.Blogs
	err := r.db.Preload("Users").First(&blog,id).Error
	if err != nil {
	  return blog,err
	}

	err = r.db.Model(&blog).Association("Tags").Clear();
	if err != nil{
		return blog,err
	}
	
	if err := r.db.Where("blog_id = ?", id).Delete(&models.Comments{}).Error; err != nil {
		return blog, err
	}
	
	helpers.DeleteFile(strings.ReplaceAll(blog.Image, "/", "\\"))

	err = r.db.Delete(&blog).Error
	if err != nil{
	return blog, err
	}

	return blog,nil
}

func (r *blogRepo) Search(keyword string, blog []models.Blogs)([]models.Blogs,error){
	err := r.db.Preload("Tags").Preload("Users").Omit("Comments").Where("title LIKE ?", "%"+keyword+"%").Find(&blog).Error
	if err != nil {
		return blog, err
	}
	return blog,nil
}