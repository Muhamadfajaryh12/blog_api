package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
	"github.com/muhamadfajaryh12/api_blogs/services"
)

type DashboardHandler struct {
	dashboardService services.DashboardService
}

func NewDashboardHandler(service services.DashboardService) * DashboardHandler{
	return &DashboardHandler{dashboardService:service}
}

// Tag godoc
// @Summary  GET DASHBOARD
// @Description GET DASHBOARD
// @Tags Dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponseSuccessDTO
// @Failure 400 {object} dto.ResponseErrorDTO
// @Failure 401 {object} dto.ResponseErrorDTO
// @Failure 500 {object} dto.ResponseErrorDTO
// @Router /dashboard [get]
func (h *DashboardHandler) Get(c *gin.Context){
	// GET user ID dari token jwt
	userIDRaw, exists := c.Get("UserID")

	if !exists {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: "User ID not found in token"})
		return
	}

	// Konversi user ID ke uint
	floatID, ok := userIDRaw.(float64)
	if !ok {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message: "Invalid user ID type"})
		return
	}

	userID := uint64(floatID)

	fmt.Println(userID)
	result, err := h.dashboardService.GetDashboard(uint64(userID))
	if err != nil {
		helpers.ErrorHandle(c, helpers.InternalServerError{Message:err.Error()})
	}

	c.JSON(http.StatusOK, dto.ResponseSuccessDTO{
		Status:http.StatusOK,
		Message:"Fetched successully",
		Data:result,
	})
}