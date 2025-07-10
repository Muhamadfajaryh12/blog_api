package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/middlewares"
)

func TagRoute(r *gin.RouterGroup, tagHandler *handlers.TagHandler) {
	tags := r.Group("/tags")
	{
		tags.GET("/",tagHandler.GetAll)
		tags.GET("/:id",tagHandler.GetById)
		
		tagsAuth := tags.Use(middlewares.Authorization())
		tagsAuth.POST("/",tagHandler.Create)
		tagsAuth.PUT("/:id",tagHandler.Update)
		tagsAuth.DELETE("/:id",tagHandler.Delete)
	}


}