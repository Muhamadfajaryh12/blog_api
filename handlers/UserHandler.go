package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler{
	return &UserHandler{Repo: repo}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.UserRequestDTO true "User registration data"
// @Success 201 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /users/register [post]
func (h *UserHandler) Register(c *gin.Context){
	var input dto.UserRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,dto.ResponseErrorDTO{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user := models.Users {
		Email: input.Email,
		Name: input.Name,
		Password: input.Password,
		Role:input.Role,
	} 

	result,err := h.Repo.Create(user)
	if err != nil{
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	response := dto.UserResponseDTO{
		ID: result.ID,
		Name: result.Name,
		Email: result.Email,
		Role: result.Role,
	}

	c.JSON(http.StatusCreated,dto.ResponseSuccessDTO{
		Status: http.StatusCreated,
		Message: "Created successfully",
		Data: response,
	})

}

// Login godoc
// @Summary User login
// @Description Authenticate user and get access token
// @Tags users
// @Accept json
// @Produce json
// @Param input body LoginRequest true "Login credentials"
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context){
	var input LoginRequest
	if err:= c.ShouldBind(&input); err != nil{
		c.JSON(http.StatusBadRequest,dto.ResponseErrorDTO{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	result,err := h.Repo.Get(input.Email,input.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized,dto.ResponseErrorDTO{
			Status: http.StatusUnauthorized,
			Message:"Email or passwort doesn't match",
		})
		return 
	}

	token,err := helpers.GenerateToken(result.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	response := dto.LoginResponseDTO{
		ID: result.ID,
		Name: result.Name,
		Role: result.Role,
		Token: token,
	}
	
	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status:http.StatusOK,
		Message:"Login successfully",
		Data:response,
	})

}