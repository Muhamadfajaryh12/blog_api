package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type BlogHandler struct {
	Repo repository.BlogRepository
}

func NewBlogHandler(repo repository.BlogRepository) * BlogHandler{
	return &BlogHandler{Repo: repo}
}

func (h *BlogHandler) Create(c *gin.Context){
	var blog models.Blogs
	err := c.ShouldBind(&blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	fileUpload, err := c.FormFile("upload")
	if err == nil{
		filePath, err := helpers.SaveFile(fileUpload,"banner")
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		blog.Image = filePath
	}

    tagIDs := c.PostFormArray("tags_id")
	var tags []models.Tags
	for _, idStr := range tagIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue 
		}
		tags = append(tags, models.Tags{ID: uint64(id)})
	}
	blog.Tags = tags
	
	result, err := h.Repo.Create(blog)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
	"message":"Created successfully",
	"data":result,
	"status":http.StatusCreated,
	})

}