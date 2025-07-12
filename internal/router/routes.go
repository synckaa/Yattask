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
	api.POST("/task", taskController.Create)
	api.POST("/taskai", taskController.CreateWithAI)
	api.PUT("/task/:id", taskController.Update)
	api.DELETE("/task/:id", taskController.Delete)

	auth := r.Group("/api")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/logout", userController.Logout)

	return r
}
