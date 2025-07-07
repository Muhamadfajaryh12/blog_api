package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type TagHandler struct {
	Repo repository.TagRepository
}

func NewTagHandler(repo repository.TagRepository) *TagHandler{
	return &TagHandler{Repo: repo}
}

func (h *TagHandler) Create(c *gin.Context){
	var tag models.Tags
	err:= c.ShouldBindJSON(&tag)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result,err := h.Repo.Create(tag)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Created successfully",
		"data":result,
		"status":http.StatusCreated,
	})
}