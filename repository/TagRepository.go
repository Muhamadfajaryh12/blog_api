package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	Create(tag models.Tags)(models.Tags,error)
	GetAll(tag models.Tags)(models.Tags,error)
}


type tagRepo struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepo{db:db}
}

func (r *tagRepo) Create(tag models.Tags)(models.Tags,error){
	err := r.db.Create(&tag).Error
	if err != nil{
		return tag,err
	}
	return tag,nil
}

func (r *tagRepo) GetAll(tag models.Tags)(models.Tags,error){
	err:= r.db.Find(&tag).Error
	if err!= nil{
		return tag,err
	}
	return tag,nil
}