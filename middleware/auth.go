package middleware

import (
	"fmt"
	"net/http"

	auth "github.com/JerryJeager/amor-rendezvous-backend/http"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := auth.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
