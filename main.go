package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-jwt-auth-boilerplate/config"
	"go-jwt-auth-boilerplate/database"
	"go-jwt-auth-boilerplate/handlers"
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
	r.POST("/register", handlers.Register(cfg))
	r.POST("/login", handlers.Login(cfg))

	fmt.Println("Server running on port", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
