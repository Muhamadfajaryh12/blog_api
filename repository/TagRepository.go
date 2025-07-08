package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	Create(tag models.Tags)(models.Tags,error)
	GetAll(tag []models.Tags)([]models.Tags,error)
	Update(id string, inputTag models.Tags)(models.Tags,error)
	Delete(id string)(models.Tags, error)
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

func (r *tagRepo) GetAll(tag []models.Tags)([]models.Tags,error){
	err:= r.db.Find(&tag).Error
	if err!= nil{
		return tag,err
	}
	return tag,nil
}

func (r *tagRepo) Update(id string, inputTag models.Tags)(models.Tags,error){
	var tag models.Tags
	err:= r.db.First(&tag,id).Error
	if err != nil{
		return tag, err
	}

	err = r.db.Model(&tag).Updates(inputTag).Error
	if err != nil{
		return tag, err
	}

	return tag,nil
}

func (r *tagRepo) Delete(id string)(models.Tags, error){
	var tag models.Tags
	err:= r.db.First(&tag,id).Error
	if err != nil{
		return tag, err
	}

	err = r.db.Delete(&tag).Error
	if err != nil{
		return tag, err
	}

	return tag, nil
}