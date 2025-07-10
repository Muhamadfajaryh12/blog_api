package models

type Comments struct {
	ID      int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	Content string `gorm:"type:varchar(255)" `
	UserID  uint64 `gorm:"index"`
	Users   *Users `gorm:"foreignKey:UserID;references:ID" `
	BlogID  uint64 `gorm:"index"`
	Blogs   *Blogs `gorm:"foreignKey:BlogID;references:ID" `
}