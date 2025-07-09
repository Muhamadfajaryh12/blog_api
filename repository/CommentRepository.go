package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type CommentRepository interface{
	Create(comment models.Comments)(models.Comments, error)
	Delete(id uint64)(models.Comments, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository{
	return &commentRepo{db:db}
}

func (r *commentRepo) Create(comment models.Comments)(models.Comments,error){
	err := r.db.Create(&comment).Error
	if err != nil{
		return comment, err
	}
	return comment, nil
}

func (r *commentRepo) Delete(id uint64)(models.Comments, error){
	var comment models.Comments
	
	err:= r.db.Find(&comment, id).Error
	if err != nil {
		return comment, err
	}

	err = r.db.Delete(&comment).Error
	if err != nil{
		return comment, err
	}

	return comment, err
}