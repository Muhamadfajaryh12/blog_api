package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
)

func TagRoute(r *gin.Engine, tagHandler *handlers.TagHandler) {
	tags := r.Group("/tags")
	{
		tags.GET("/",tagHandler.GetAll)
		tags.GET("/:id",tagHandler.GetById)
		tags.POST("/",tagHandler.Create)
		tags.PUT("/:id",tagHandler.Update)
		tags.DELETE("/:id",tagHandler.Delete)
	}

}