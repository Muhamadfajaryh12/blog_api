package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/mapper"
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

func (h *BlogHandler) GetAll(c *gin.Context){
	var blogs []models.Blogs
	result, err := h.Repo.GetAll(blogs)

		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}

	var response []dto.BlogResponseDTO
	for _, blog := range result {
		response = append(response, mapper.BlogResponse(blog))
	}
		c.JSON(http.StatusOK,gin.H{
			"data":response,
			"status":http.StatusOK,
		})
}

func (h *BlogHandler) GetDetail(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var blog models.Blogs

	result, err := h.Repo.GetDetail(uint64(id), blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	response := mapper.BlogDetailResponse(result)
	c.JSON(http.StatusOK,gin.H{
		"data":response,
		"status":http.StatusOK,
	})
}

func (h *BlogHandler) Update(c *gin.Context){
	ParamId := c.Param("id")
	var blog models.Blogs
	
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	err = c.ShouldBind(&blog)
	if err != nil {
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
	for _, idStr := range tagIDs {
		if tagId, err := strconv.Atoi(idStr); err == nil {
			blog.Tags = append(blog.Tags, models.Tags{ID: uint64(tagId)})
		}
	}
	
	result, err := h.Repo.Update(uint64(id), blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
	"message":"Updated successfully",
	"data":result,
	"status":http.StatusOK,
	})
}

func (h *BlogHandler) Delete(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	result,err := h.Repo.Delete(uint64(id))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
	"message":"Deleted successfully",
	"data":result,
	"status":http.StatusOK,
	})
}