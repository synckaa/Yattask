package router

import (
	"Yattask/internal/controller/taskcontrollers"
	"Yattask/internal/controller/usercontrollers"
	"Yattask/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AllRoutes(userController usercontrollers.UserController, taskController taskcontrollers.TaskController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api", middleware.AuthMiddleware)
	api.GET("/dashboard", userController.GetProfile)
	api.POST("/tasks", taskController.Create)
	api.POST("/tasks/ai", taskController.CreateWithAI)
	api.PUT("/tasks/:id", taskController.Update)
	api.DELETE("/tasks/:id", taskController.Delete)
	api.GET("/tasks/:id", taskController.GetByIDWithTask)

	auth := r.Group("/api/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/logout", userController.Logout)

	return r
}
