package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/dto"
	"github.com/muhamadfajaryh12/api_blogs/helpers"
)

func Authorization() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader,"Bearer ")

		claims,err := helpers.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseErrorDTO{
				Status:http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		c.Set("UserID",claims["user_id"])
		c.Next()
	}
}