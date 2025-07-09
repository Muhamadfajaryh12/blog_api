package dto

type UserDTO struct {
	ID       uint64 `json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

type UserResponseDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}