// This package handles all the routes (URLs) in our application
package routes

// Import necessary packages
import (
	"go-gin-auth/controllers" // Our route handlers
	"go-gin-auth/middleware"  // Our middleware

	"github.com/gin-gonic/gin" // Web framework
)

// SetupRoutes configures all the URLs our application will respond to
// router: The Gin engine that will handle all web requests
func SetupRoutes(router *gin.Engine) {
	// Create a versioned API group
	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/dashboard", middleware.AuthMiddleware(), controllers.Dashboard)
		v1.GET("/users", middleware.AuthMiddleware(), controllers.ListUsers)

		// Group all post routes and apply AuthMiddleware once
		postRoutes := v1.Group("/posts", middleware.AuthMiddleware())
		{
			postRoutes.GET("", controllers.ListPosts)
			postRoutes.POST("", controllers.CreatePost)
			postRoutes.GET("/:id", controllers.GetPost)
			postRoutes.PUT("/:id", controllers.UpdatePost)
			postRoutes.DELETE("/:id", controllers.DeletePost)
		}
	}
}
