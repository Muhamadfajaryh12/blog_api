package models

import "time"

type Blogs struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" `
	Title     string `gorm:"type:varchar(255)"`
	Content   string `gorm:"type:varchar(255)"`
	Image     string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UserID   uint       `gorm:"index"`
	Users    *Users     `gorm:"foreignKey:UserID"`
	Tags     []Tags     `gorm:"many2many:blog_tags;"`
	Comments []Comments `gorm:"constraint:OnDelete:CASCADE;foreignKey:BlogID;references:ID"`
	Views []ViewBlog `gorm:"contraint:OnDeleete:CASCADE;foreignKey:BlogID;references:ID"`
}