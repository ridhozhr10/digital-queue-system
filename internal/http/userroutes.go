package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOrigins: []string{
			"*",
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "User service is running!", "header": ctx.Request.Header})
		})
	}
}
