package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-jwt-auth-boilerplate/config"
	"go-jwt-auth-boilerplate/database"
	"go-jwt-auth-boilerplate/handlers"
	"go-jwt-auth-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	database.Connect(cfg)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		auth.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user_id": c.MustGet("userID"),
				"email":   c.MustGet("email"),
			})
		})
	}
	// Rate limit: max 5 requests per minute on auth routes
	r.POST("/register", middleware.RateLimit(5, time.Minute), handlers.Register(cfg))
	r.POST("/login", middleware.RateLimit(5, time.Minute), handlers.Login(cfg))
	// r.POST("/register", handlers.Register(cfg))
	// r.POST("/login", handlers.Login(cfg))

	fmt.Println("Server running on port", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
