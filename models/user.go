package models

type Users struct {
	ID       uint64     `gorm:"primaryKey;autoIncrement" `
	Name     string     `gorm:"type:varchar(255)"  `
	Email    string     `gorm:"type:varchar(255)" `
	Password string     `gorm:"type:varchar(255)" `
	Role     string     `gorm:"type:varchar(255)"  `
	Blogs    []Blogs    `gorm:"foreignKey:UserID" `
	Comments []Comments `gorm:"foreignKey:UserID"`
}
