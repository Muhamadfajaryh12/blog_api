package dto

type UserRequestDTO struct {
	Name     string `form:"name" json:"name" example:"test" binding:"required"`
	Email    string `form:"email" json:"email" example:"email@gmail.com" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" example:"admin" binding:"required"`
}

type UserResponseDTO struct {
	ID    uint64 `json:"id" example:"1"`
	Name  string `json:"name" example:"test"`
	Email string `json:"email" example:"email@gmail.com"`
	Role  string `json:"role" example:"admin"`
}

type LoginResponseDTO struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" example:"test"`
	Role  string `json:"role" example:"admin"`
	Token string `json:"token" exampl:"123"`
}