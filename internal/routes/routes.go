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
	propertyHandler := handlers.NewPropertyHandler()
	utilityHandler := handlers.NewUtilityHandler()
	postHandler := handlers.NewPostHandler()
	postFavoriteHandler := handlers.NewPostFavoriteHandler()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.PUT("/change-password", middlewares.AuthMiddleware(), authHandler.ChangePassword)
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
		property := api.Group("/property")
		{
			property.GET("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord)), propertyHandler.GetByID)
			property.POST("/search", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord), string(constants.AccountRoleTenant)), propertyHandler.Search)
			property.POST("", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord)), propertyHandler.Create)
			property.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord)), propertyHandler.Update)
			property.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord)), propertyHandler.Delete)
		}
		utility := api.Group("/utility")
		{
			utility.GET("", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin), string(constants.AccountRoleLandlord)), utilityHandler.GetAll)
			utility.POST("", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), utilityHandler.Create)
			utility.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), utilityHandler.Update)
			utility.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleAdmin)), utilityHandler.Delete)
		}
		postFavorite := api.Group("/post-favorite", middlewares.AuthMiddleware(), middlewares.RequireRoles(string(constants.AccountRoleTenant)))
		{
			postFavorite.GET("", postFavoriteHandler.Search)
			postFavorite.GET("/ids", postFavoriteHandler.IDs)
			postFavorite.POST("/:postId", postFavoriteHandler.Add)
			postFavorite.DELETE("/:postId", postFavoriteHandler.Remove)
		}
		post := api.Group("/post", middlewares.AuthMiddleware())
		{
			post.GET("/:id", middlewares.RequireRoles(string(constants.AccountRoleLandlord), string(constants.AccountRoleTenant)), postHandler.GetByID)
			post.POST("/search", middlewares.RequireRoles(string(constants.AccountRoleLandlord)), postHandler.Search)
			post.POST("", middlewares.RequireRoles(string(constants.AccountRoleLandlord)), postHandler.Create)
			post.PUT("/:id", middlewares.RequireRoles(string(constants.AccountRoleLandlord)), postHandler.Update)
			post.DELETE("/:id", middlewares.RequireRoles(string(constants.AccountRoleLandlord)), postHandler.Delete)
		}
	}
}
