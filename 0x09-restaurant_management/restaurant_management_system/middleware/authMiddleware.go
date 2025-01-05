package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "restaurant_management_system/helpers"
)


func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
            // c.AbortWithStatusJSON(401, gin.H{"error": "unauthorised"})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorised"})
			c.Abort()
            return
        }

		claims, err := helpers.ValidateToken(clientToken)
		if err!="" {
			// c.AbortWithStatusJSON(401, gin.H{"error": err})
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            c.Abort()
            return
        }
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}