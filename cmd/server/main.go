package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"auth-backend/internal/db"
	"auth-backend/internal/handlers"
)

func main() {
	database, err := db.Init("./auth.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()
	authHandler := handlers.NewAuthHandler(database)

	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server launch failed: %v", err)
	}
}
