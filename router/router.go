package router

import (
	"golang-starter-kit/handlers"
	"golang-starter-kit/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{

		userRoutes := v1.Group("user")
		{
			userRoutes.POST("/new", middleware.Cors, handlers.CreateUser)
			userRoutes.GET("/show/:userId", middleware.Cors, handlers.GetUser)
			userRoutes.PUT("/:userId", middleware.Cors, handlers.UpdateUser)
			userRoutes.DELETE("/:userId", middleware.Cors, handlers.DeleteUser)
		}
		v1.GET("/users", middleware.Cors, handlers.GetUsers)
	}

}
