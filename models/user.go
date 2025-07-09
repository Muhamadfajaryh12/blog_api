package models

type Users struct {
	ID       uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string     `gorm:"type:varchar(255)" json:"name" `
	Email    string     `gorm:"type:varchar(255)" json:"email" `
	Password string     `gorm:"type:varchar(255)" json:"password" `
	Role     string     `gorm:"type:varchar(255)" json:"role" `
	Blogs    []Blogs    `gorm:"foreignKey:UserID" json:"blogs"`
	Comments []Comments `gorm:"foreignKey:UserID" json:"comments"`
}
