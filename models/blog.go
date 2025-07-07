package models

type Blogs struct {
	ID      uint64 `gorm:"primarKey;autoIncrement" json:"id"`
	Title   string `gorm:"type:varchar(255)" form:"title" json:"title" validate:"required"`
	Content string `gorm:"type:varchar(255)" form:"contant" json:"content" validate:"required"`
	Image   string `gorm:"type:varchar(255)" form:"image" json:"image" validate:"required"`
	UserID  uint   `gorm:"index"`
	Users   Users  `gorm:"foreignKey:UserID" json:"author"`
	Tags    []Tags `gorm:"many2many:blog_tags;" json:"tags"`
}