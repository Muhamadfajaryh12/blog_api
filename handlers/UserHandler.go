package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler{
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) Register(c *gin.Context){
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result,err := h.Repo.Create(user)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"status":http.StatusCreated,
		"data":result,
		"message":"Register successfully",
	})
}

func (h *UserHandler) Login(c *gin.Context){
	var input LoginRequest
	if err:= c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result,err := h.Repo.Get(input.Email,input.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"Email or password doesn't match",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"data":result,
		"message":"",
	})

}