package models

type Comments struct {
	ID      int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	Content string `gorm:"type:varchar(255)" json:"content" form:"content" `
	UserID  uint64 `gorm:"index"  json:"user_id"`
	Users   *Users `gorm:"foreignKey:UserID;references:ID" json:"-"`
	BlogID  uint64 `gorm:"index" json:"blog_id"`
	Blogs   *Blogs `gorm:"foreignKey:BlogID;references:ID" json:"-"`
}