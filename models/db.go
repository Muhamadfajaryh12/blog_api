package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectionDatabase() *gorm.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)
	

	DB, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("error",err)
	}

	DB.AutoMigrate(&Users{},&Blogs{},&Tags{},&Comments{})
	return DB
}