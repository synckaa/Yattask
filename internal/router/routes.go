package router

import (
	"Yattask/internal/controller/taskcontrollers"
	"Yattask/internal/controller/usercontrollers"
	"Yattask/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AllRoutes(userController usercontrollers.UserController, taskController taskcontrollers.TaskController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api", middleware.AuthMiddleware)
	api.GET("/dashboard", userController.GetProfile)
	api.POST("/tasks", taskController.Create)
	api.POST("/ai/tasks", taskController.CreateWithAI)
	api.PUT("/tasks/:id", taskController.Update)
	api.DELETE("/tasks/:id", taskController.Delete)

	auth := r.Group("/api/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/logout", userController.Logout)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
