package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
)

func BlogRoutes(r *gin.Engine, blogHandler *handlers.BlogHandler){
	blogs := r.Group("/blogs")
	{
		blogs.GET("/")
		blogs.POST("/",blogHandler.Create)
	}
}