package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/middlewares"
)

func DashboardRoute(r *gin.RouterGroup, dashboardHandler *handlers.DashboardHandler){
	dashboard := r.Group("/dashboard")
	dashboardAuth := dashboard.Use(middlewares.Authorization())
	{
		dashboardAuth.GET("",dashboardHandler.Get)		
	}
}