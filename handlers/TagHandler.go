package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/mapper"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
)

type TagHandler struct {
	Repo repository.TagRepository
}

func NewTagHandler(repo repository.TagRepository) *TagHandler{
	return &TagHandler{Repo: repo}
}

// Tag godoc
// @Summary Create new Tag
// @Description Add a new tag
// @Tags tags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.TagDTO true "Tag Data"
// @Success 201 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /tags [post]
func (h *TagHandler) Create(c *gin.Context){
	var input dto.TagDTO
	err:= c.ShouldBind(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest,dto.ResponseErrorDTO{
			Status:http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	tag:= models.Tags{
		Tag: input.Tag,
	}
	
	result,err := h.Repo.Create(tag)
	if err != nil{
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,dto.ResponseSuccessDTO{
		Status: http.StatusCreated,
		Message: "Created succesfully",
		Data: result,
	})
}

// Tag godoc
// @Summary Get All Tag
// @Description Get all the Tag
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /tags [get]
func (h *TagHandler) GetAll(c *gin.Context){
	var tag []models.Tags
	result, err := h.Repo.GetAll(tag)


	if err != nil{
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var response []dto.TagResponseDTO
	for _,t := range result {
		response = append(response, mapper.TagRespose(t))
	}

	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Fetched succesfully",
		Data: response,
	})
}

// Tag godoc
// @Summary GET DETAIL TAG
// @Description Get Detail the Tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID" Example(1)
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /tags/{id} [get]
func (h *TagHandler) GetById(c *gin.Context){
	ParamId := c.Param("id")
	id, err := strconv.Atoi(ParamId)

	if err != nil {
		c.JSON(http.StatusBadRequest,dto.ResponseErrorDTO{
			Status:http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var tag models.Tags

	result,err := h.Repo.GetById(uint64(id),tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErrorDTO{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	response := mapper.TagDetailResponse(result)

	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Fetched succesfully",
		Data: response,
	})
}

// Tag godoc
// @Summary UPDATE TAG
// @Description UPDATE the Tag
// @Tags tags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Tag ID" Format(int64) Example(1)
// @Param request body dto.TagDTO true "Tag update data"
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /tags/{id} [put]
func (h *TagHandler) Update(c *gin.Context){
	id := c.Param("id")
	
	var inputTag models.Tags
	if err := c.ShouldBind(&inputTag);
	err != nil{
		c.JSON(http.StatusBadRequest,dto.ResponseErrorDTO{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	result,err := h.Repo.Update(id, inputTag)
	if err != nil{
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Updated successfully",
		Data: result,
	})
}

// Tag godoc
// @Summary DELETE TAG
// @Description DELETE the Tag
// @Tags tags
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Tag ID" Example(1)
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /tags/{id} [delete]
func (h *TagHandler) Delete(c *gin.Context){
	id := c.Param("id")

	result,err := h.Repo.Delete(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError,dto.ResponseErrorDTO{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccessDTO{
		Status: http.StatusOK,
		Message: "Deleted successfully",
		Data: result,
	})
}