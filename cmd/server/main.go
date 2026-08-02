package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"easyrent-server/internal/database"
	"easyrent-server/internal/routes"
	"easyrent-server/internal/services"
	"easyrent-server/internal/workers"
)

func main() {
	database.Connect()

	emailService := services.NewEmailService()

	workers.StartEmailWorker(emailService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/storage", "./storage")

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
