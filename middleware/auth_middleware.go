package middleware

import (
	"TFP/utils"
	helper "TFP/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			response := utils.BuildAuthErrorResponse()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenString := cookie.Value

		claims, msg := helper.ValidateToken(tokenString)
		if msg != "" {
			response := utils.BuildAuthErrorResponse()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}