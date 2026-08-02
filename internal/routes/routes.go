package routes

import (
	"github.com/gin-gonic/gin"

	"easyrent-server/internal/constants"
	"easyrent-server/internal/handlers"
	"easyrent-server/internal/middlewares"
)

// RegisterRoutes sets up the API routes for the application, including a health check endpoint and authentication routes.
func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authHandler := handlers.NewAuthHandler()
	userHandler := handlers.NewUserHandler()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
		user := api.Group("/user")
		{
			user.GET("/me", middlewares.AuthMiddleware(), userHandler.GetMe)
			user.PUT("/me", middlewares.AuthMiddleware(), userHandler.UpdateMe)
			user.POST("/search", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), userHandler.Search)
			user.GET("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), userHandler.GetByID)
			user.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), userHandler.Update)
			user.POST("", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), userHandler.Create)
		}
	}
}
