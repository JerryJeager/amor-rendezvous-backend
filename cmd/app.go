package cmd

import (
	"github.com/JerryJeager/amor-rendezvous-backend/api"
	"github.com/JerryJeager/amor-rendezvous-backend/manualwire"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()

func ExecuteApiRoutes() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})

	v1 := r.Group("/api/v1")

	v1.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, api.OpenApiDocs())
	})

	users := v1.Group("/users")
	users.POST("", userController.CreateUser)

	r.Run(":8080")
}