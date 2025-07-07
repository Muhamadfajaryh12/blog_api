package models

import "time"

type Tags struct {
	ID        uint64  `gorm:"primaryKey; autoIncrement" json:"id"`
	Tag       string `gorm:"type:varchar(255)" form:"tag" json:"tag" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	Blogs []Blogs `gorm:"many2many:blog_tags;" json:"blogs"`
}