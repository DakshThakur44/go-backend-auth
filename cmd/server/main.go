package main

import (
	"log"
	"net/http"
	"time"

	"auth-backend/internal/db"
	"auth-backend/internal/handlers"
	"auth-backend/internal/middleware"

	"github.com/alitto/pond"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Init("./auth.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	pool := pond.New(10, 10)

	r := gin.Default()
	authHandler := handlers.NewAuthHandler(database, pool)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "yo wsg"})
	})

	lmt := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{
		DefaultExpirationTTL: 1 * time.Hour,
	})

	lmt.SetMessage(`{"error": "Too many requests, please try again later."}`)
	lmt.SetMessageContentType("application/json")

	reqQueue := middleware.LimitConcurrency(2, 10*time.Second)

	v1 := r.Group("/api/v1/auth")
	v1.Use(tollbooth_gin.LimitHandler(lmt))
	v1.Use(reqQueue)
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server launch failed: %v", err)
	}
}
