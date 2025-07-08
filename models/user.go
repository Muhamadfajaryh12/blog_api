package models

type Users struct {
	ID       uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string     `gorm:"type:varchar(255)" form:"name" json:"name" binding:"required"`
	Email    string     `gorm:"type:varchar(255)" form:"email" json:"email" binding:"required"`
	Password string     `gorm:"type:varchar(255)" form:"password" json:"password" binding:"required"`
	Role     string     `gorm:"type:varchar(255)" form:"role" json:"role" binding:"required"`
	Blogs    []Blogs    `gorm:"foreignKey:UserID" json:"blogs"`
	Comments []Comments `gorm:"foreignKey:UserID" json:"comments"`
}
