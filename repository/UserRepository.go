package repository

import (
	"github.com/muhamadfajaryh12/api_blogs/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.Users)(models.Users,error)
	Get(email string, password string)(models.Users,error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepo{db:db}
}


func (r *userRepo) Create(user models.Users)(models.Users,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)
	err = r.db.Create(&user).Error
	return user,err
}

func (r *userRepo) Get(email string, password string) (models.Users, error){
	var user models.Users
	if err := r.db.Where("email = ?",email).First(&user).Error;
	err != nil{
		return user,err
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); 
	err != nil{
		return user,err
	}

	return user,nil
}