package routes

import (
	"go-task-manager-mvc/controllers"
	"go-task-manager-mvc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// 🔐 Rutas públicas
	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)

	// 🔒 Rutas protegidas con JWT
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", controllers.GetTasks)
		protected.POST("/tasks", controllers.CreateTask)
		protected.PUT("/tasks/:id", controllers.UpdateTask)
		protected.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}
