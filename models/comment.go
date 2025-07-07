package models

type Comments struct {
	ID      int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	Content string `gorm:"type:varchar(255)" json:"content" form:"content" validate:"required"`
	UserID  uint64 `gorm:"index"  json:"user_id" form:"user_id" validate:"required"`
	User    Users  `gorm:"foreignKey:UserID" json:"user"`
	BlogID  uint64 `gorm:"index" json:"blog_id" form:"blog_id" validate:"required"`
	Blog    Blogs  `gorm:"foreignKey:BlogID" json:"blog"`
}