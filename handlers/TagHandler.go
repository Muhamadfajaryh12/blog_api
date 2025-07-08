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
	err:= c.ShouldBind(&tag)
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

func (h *TagHandler) GetAll(c *gin.Context){
	var tag []models.Tags
	result, err := h.Repo.GetAll(tag)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{
		"data":result,
		"status":http.StatusOK,
	})
}

func (h *TagHandler) Update(c *gin.Context){
	id := c.Param("id")
	
	var inputTag models.Tags
	if err := c.ShouldBind(&inputTag);
	err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result,err := h.Repo.Update(id, inputTag)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Updated successfully",
		"data":result,
		"status":http.StatusOK,
	})
}

func (h *TagHandler) Delete(c *gin.Context){
	id := c.Param("id")

	result,err := h.Repo.Delete(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Deleted successfully",
		"data":result,
		"status":http.StatusOK,
	})
}