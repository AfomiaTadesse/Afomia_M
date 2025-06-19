package router

import (
	"github.com/AfomiaTadesse/Afomia_M/backend/controller"
	"github.com/AfomiaTadesse/Afomia_M/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userCtrl *controller.UserController,
	movieCtrl *controller.MovieController,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		// User routes (no auth required)
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/signup", userCtrl.Signup)
			userRoutes.POST("/login", userCtrl.Login)
		}

		// Movie routes (auth required)
		movieRoutes := api.Group("/movies")
		movieRoutes.Use(middleware.AuthMiddleware())
		{
			movieRoutes.POST("/", movieCtrl.CreateMovie)
			movieRoutes.GET("/", movieCtrl.GetMovies)
			movieRoutes.GET("/search", movieCtrl.SearchMovies)
			movieRoutes.GET("/:id", movieCtrl.GetMovieByID)
			movieRoutes.PUT("/:id", movieCtrl.UpdateMovie)
			movieRoutes.DELETE("/:id", movieCtrl.DeleteMovie)
		}
	}

	return router
}