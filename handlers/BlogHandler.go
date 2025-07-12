package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

// Tag godoc
// @Summary Create Blog
// @Description create new a blog
// @Tags Blogs
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Blog title"
// @Param content formData string true "Blog content"
// @Param upload formData file true "Blog image file"
// @Param user_id formData int true "Author"
// @Param tags_id formData []int true "Array of Tag IDs"
// @Success 201 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /blogs [post]
func (h *BlogHandler) Create(c *gin.Context){
	var input dto.BlogRequestDTO
	var filePath string
	err := c.ShouldBind(&input)
	if err != nil{
		helpers.ErrorHandle(c, helpers.BadRequestError{Message: err.Error()})
		return
	}

	fileUpload, err := c.FormFile("upload")
	if err == nil{
		filePath, err = helpers.SaveFile(fileUpload,"banner")
		if err != nil{
			helpers.ErrorHandle(c, helpers.BadRequestError{Message: err.Error()})
			return
		}
	}

    rawTags := c.PostForm("tags_id")
	tagIDs := strings.Split(rawTags, ",")
	var tags []models.Tags
	for _, idStr := range tagIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue 
		}
		tags = append(tags, models.Tags{ID: uint64(id)})
	}

	blog := models.Blogs{
		Title: input.Title,
		Content: input.Content,
		Image:  filePath,
		UserID: uint(input.UserID),
		Tags:tags,
	}
	
	result, err := h.Repo.Create(blog)
	if err != nil{
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	response := mapper.BlogResponse(result)
	c.JSON(http.StatusCreated, dto.ResponseSuccessDTO{
		Status: http.StatusCreated,
		Message: "Created successfully",
		Data: response,
	})
}

// Tag godoc
// @Summary Get All Blog
// @Description Get all the blog
// @Tags Blogs
// @Produce json
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /blogs [get]
func (h *BlogHandler) GetAll(c *gin.Context){
	var blogs []models.Blogs
	getAll, err := h.Repo.GetAll(blogs)
	if err != nil{
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	getTrending, err := h.Repo.GetTrending(blogs)
	if err != nil{
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	
	getLatest, err := h.Repo.GetLatest(blogs)
	if err != nil{
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}
	var
	 (
		allResponse []dto.BlogResponseDTO 
		trendingResponse []dto.BlogResponseDTO
		latestResponse []dto.BlogResponseDTO
	)

	for _, blog := range getAll {
		allResponse = append(allResponse, mapper.BlogResponse(blog))
	}
	
	for _, blog := range getTrending {
		trendingResponse = append(trendingResponse, mapper.BlogResponse(blog))
	}

		for _, blog := range getLatest {
		latestResponse = append(latestResponse, mapper.BlogResponse(blog))
	}
	
	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Fetched successfully",
		Data:gin.H{
			"all": allResponse,
			"trending":trendingResponse,
			"latest":latestResponse,
		},
	})
}



// Tag godoc
// @Summary Get Detail Blog
// @Description Get detail a blog
// @Tags Blogs
// @Produce json
// @Param id path int true "Blog ID" Format(int64) Example(1)
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /blogs/{id} [get]
func (h *BlogHandler) GetDetail(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message:"Invalid ID"})
		return
	}
	var blog models.Blogs

	result, err := h.Repo.GetDetail(uint64(id), blog)
	if err != nil {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	response := mapper.BlogDetailResponse(result)
	c.JSON(http.StatusOK,gin.H{
		"data":response,
		"status":http.StatusOK,
	})
}

// Tag godoc
// @Summary Update Blog
// @Description update new a blog
// @Tags Blogs
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Blog title"
// @Param content formData string true "Blog content"
// @Param upload formData file optional "Blog image file"
// @Param user_id formData int true "author"
// @Param tags_id formData []int true "Array of Tag IDs"
// @Param id path int true "Blog ID"
// @Success 201 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /blogs/{id} [put]
func (h *BlogHandler) Update(c *gin.Context){
	ParamId := c.Param("id")
	var blog models.Blogs
	var input dto.BlogRequestDTO
	var filePath string
	var tags []models.Tags

	id, err := strconv.Atoi(ParamId)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message: "Invalid ID"})
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message:err.Error()})
		return
	}

	fileUpload, err := c.FormFile("upload")
		if err == nil{
		filePath, err := helpers.SaveFile(fileUpload,"banner")
		if err != nil{
			helpers.ErrorHandle(c, helpers.BadRequestError{Message:err.Error()})
			return
		}
		blog.Image = filePath
	}

    rawTags := c.PostForm("tags_id")
	tagIDs := strings.Split(rawTags, ",")
	for _, idStr := range tagIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue 
		}
		tags = append(tags, models.Tags{ID: uint64(id)})
	}

	blog = models.Blogs{
		Title: input.Title,
		Content: input.Content,
		Image: filePath,
		Tags:tags,
		UserID: uint(input.UserID),
	}

	result, err := h.Repo.Update(uint64(id), blog)
	if err != nil {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message:err.Error()})
		return
	}
	
	response := mapper.BlogResponse(result)
	fmt.Printf("RESULT: %+v\n", result)

	c.JSON(http.StatusOK, dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Updated successfully",
		Data: response,
	})
}


// Tag godoc
// @Summary Delete Blog
// @Description Delete a blog
// @Tags Blogs
// @Produce json
// @Security BearerAuth
// @Param id path int true "Blog ID" Format(int64) Example(1)
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /blogs/{id} [delete]
func (h *BlogHandler) Delete(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		helpers.ErrorHandle(c, helpers.BadRequestError{Message: "Invalid ID"})
		return
	}

	result,err := h.Repo.Delete(uint64(id))
	if err != nil{
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: err.Error()})
		return
	}

	response := mapper.BlogDetailResponse(result)
	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status:http.StatusOK,
		Message:"Deleted successfully",
		Data: response,
	})
}