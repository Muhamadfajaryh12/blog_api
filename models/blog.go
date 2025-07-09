package models

type Blogs struct {
	ID       uint64     `gorm:"primarKey;autoIncrement" json:"id"`
	Title    string     `gorm:"type:varchar(255)" form:"title" json:"title" binding:"required"`
	Content  string     `gorm:"type:varchar(255)" form:"content" json:"content" binding:"required"`
	Image    string     `gorm:"type:varchar(255)"  json:"image"`
	UserID   uint       `gorm:"index" form:"user_id" binding:"required"`
	Users    *Users     `gorm:"foreignKey:UserID" json:"author" binding:"-"`
	Tags     []Tags     `gorm:"many2many:blog_tags;" json:"tags"`
	Comments []Comments `gorm:"foreignKey:BlogID;references:ID" json:"comments"`
}