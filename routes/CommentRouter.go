package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
)

func CommentRouter(r *gin.Engine, commentHandler *handlers.CommentHandler ){
	comments := r.Group("comments")
	{
		comments.POST("/", commentHandler.Create)
		comments.DELETE("/:id",commentHandler.Delete)
	}
}