package router

import (
	"Yattask/internal/controller/taskControllers"
	"Yattask/internal/controller/userControllers"
	"Yattask/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AllRoutes(userController userControllers.UserController, taskController taskControllers.TaskController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api", middleware.AuthMiddleware)
	api.GET("/dashboard", userController.GetProfile)
	api.POST("/tasks", taskController.Create)
	api.POST("/ai/task", taskController.CreateWithAI)
	api.PUT("/tasks/:id", taskController.Update)
	api.DELETE("/tasks/:id", taskController.Delete)

	auth := r.Group("/api/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/logout", userController.Logout)

	return r
}
