package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
)

func Authorization() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader,"Bearer ")

		claims,err := helpers.VerifyToken(tokenString)
		if err != nil {
			helpers.ErrorHandle(c, helpers.UnauthorizedError{Message: "Unauthorized"})
			c.Abort()
			return 
		}

		c.Set("UserID",claims["user_id"])
		c.Next()
	}
}