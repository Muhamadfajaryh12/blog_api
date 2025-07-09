package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type CommentHandler struct {
	Repo repository.CommentRepository
}

func NewCommentHandler(repo repository.CommentRepository) * CommentHandler{
	return &CommentHandler{Repo:repo}
}

func (h *CommentHandler) Create(c *gin.Context){
	var input dto.CommentDTO

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	comment := models.Comments{
		Content: input.Content,
		UserID:  input.UserID,
		BlogID:  input.BlogID,
	}

	result,err := h.Repo.Create(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})		
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Created successfully",
		"data":result,
		"status":http.StatusCreated,
	})
}

func (h *CommentHandler) Delete(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result,err := h.Repo.Delete(uint64(id))
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Deleted successfully",
		"data":result,
		"status":http.StatusOK,
	})
}