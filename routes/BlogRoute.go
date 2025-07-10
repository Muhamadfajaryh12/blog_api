package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/middlewares"
)

func BlogRoutes(r *gin.RouterGroup, blogHandler *handlers.BlogHandler){
	blogs := r.Group("/blogs")
	{
		blogs.GET("/",blogHandler.GetAll)
		blogs.GET("/:id",blogHandler.GetDetail)
		
		blogsAuth := blogs.Use(middlewares.Authorization())
		blogsAuth.POST("/",blogHandler.Create)
		blogsAuth.PUT("/:id",blogHandler.Update)
		blogsAuth.DELETE("/:id",blogHandler.Delete)
	}
}