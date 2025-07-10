package models

import "time"

type Tags struct {
	ID        uint64  `gorm:"primaryKey; autoIncrement" `
	Tag       string `gorm:"type:varchar(255)"`
	CreatedAt time.Time 
	Blogs []Blogs `gorm:"many2many:blog_tags;" `
}