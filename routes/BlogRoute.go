package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
)

func BlogRoutes(r *gin.Engine, blogHandler *handlers.BlogHandler){
	blogs := r.Group("/blogs")
	{
		blogs.GET("/",blogHandler.GetAll)
		blogs.GET("/:id",blogHandler.GetDetail)
		blogs.POST("/",blogHandler.Create)
		blogs.PUT("/:id",blogHandler.Update)
		blogs.DELETE("/:id",blogHandler.Delete)
	}
}