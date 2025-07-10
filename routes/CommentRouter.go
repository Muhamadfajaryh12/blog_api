package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/middlewares"
)

func CommentRouter(r *gin.RouterGroup, commentHandler *handlers.CommentHandler ){
	comments := r.Group("comments")
	comments.Use(middlewares.Authorization())
	{
		comments.POST("/", commentHandler.Create)
		comments.DELETE("/:id",commentHandler.Delete)
	}
}