package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectionDatabase() *gorm.DB{
	user := "root"
	password := ""
	database := "blog_db"
	host := "127.0.0.1"
	port := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)
	
	var err error
	DB, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("error",err)
	}

	// DB.AutoMigrate(&Users{},&Blogs{},&Tags{},&Comments{})
	return DB
}