package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
)

func UserRoutes(r *gin.Engine, userHandler *handlers.UserHandler){
	users:= r.Group("/users")
	{
		users.POST("/register",userHandler.Register)
		users.POST("/login",userHandler.Login)
	}
}