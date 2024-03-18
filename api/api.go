package api

import "github.com/gin-gonic/gin"

func ExecuteApiRoutes() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})

	r.Run(":8080")
}