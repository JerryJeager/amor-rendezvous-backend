package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/JerryJeager/amor-rendezvous-backend/api"
	"github.com/JerryJeager/amor-rendezvous-backend/manualwire"
	"github.com/JerryJeager/amor-rendezvous-backend/middleware"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()
var weddingController = manualwire.GetWeddingController()
var inviteeController = manualwire.GetInviteeController()

func ExecuteApiRoutes() {
	fmt.Println("executing api routes")

	r := gin.Default()
	r.Use(cors.Default())
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
	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.CreateToken)

	user := users
	user.Use(middleware.JwtAuthMiddleware())
	user.GET("/:user-id", userController.GetUser)

	wedding := v1.Group("/wedding")
	wedding.Use(middleware.JwtAuthMiddleware())
	{
		wedding.POST("", weddingController.CreateWedding)
		wedding.GET("/:wedding-id", weddingController.GetWedding)

		weddingEvent := wedding.Group("/:wedding-id/event-type")
		weddingEvent.POST("", weddingController.CreateEventType)
		weddingEvent.PUT("/:event-type-id", weddingController.UpdateEventType)
		weddingEvent.DELETE("/:event-type-id", weddingController.DeleteEventType)
	}

	invitation := v1.Group("/invitation")
	invitation.Use(middleware.JwtAuthMiddleware())
	{
		invitation.POST("/guest", inviteeController.CreateInvitee)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
