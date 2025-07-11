package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type CommentHandler struct {
	Repo repository.CommentRepository
}

func NewCommentHandler(repo repository.CommentRepository) * CommentHandler{
	return &CommentHandler{Repo:repo}
}

// Tag godoc
// @Summary Create Comment
// @Description create a new a comment
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CommentRequestDTO true "Comment Data"
// @Success 201 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /comments [post]
func (h *CommentHandler) Create(c *gin.Context){
	var input dto.CommentRequestDTO

	err := c.ShouldBind(&input)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message:err.Error()})
		return
	}

	comment := models.Comments{
		Content: input.Content,
		UserID:  input.UserID,
		BlogID:  input.BlogID,
	}

	result,err := h.Repo.Create(comment)
	if err != nil {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message:err.Error()})
		return
	}

	c.JSON(http.StatusCreated,dto.ResponseSuccessDTO{
		Status: http.StatusCreated,
		Message: "Created successfully",
		Data: result,
	})
}

// Tag godoc
// @Summary Create Comment
// @Description create a new a comment
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Comment ID" Format(int64) Example(1)
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /comments/{id} [delete]
func (h *CommentHandler) Delete(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message:"Invalid ID"})
		return
	}

	result,err := h.Repo.Delete(uint64(id))
	if err != nil {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Deleted successfully",
		Data: result,
	})
}