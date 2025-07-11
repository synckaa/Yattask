package main

import (
	"Yattask/configs"
	"Yattask/internal/controller/taskControllers"
	"Yattask/internal/controller/userControllers"
	"Yattask/internal/middleware"
	"Yattask/internal/repository/tagrepositories"
	"Yattask/internal/repository/taskrepositories"
	"Yattask/internal/repository/userrepositories"
	"Yattask/internal/service/taskservices"
	"Yattask/internal/service/userservices"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	configs.LoadEnv()
	configs.GetConnDB()
	configs.SyncTables(configs.DB)
}

func main() {
	validate := validator.New()
	repo := userrepositories.NewUserRepository()
	Service := userservices.NewUserService(configs.DB, repo, validate)
	Controller := userControllers.NewUserController(Service)
	tagRepository := tagrepositories.NewTagRepository()
	taskRepo := taskrepositories.NewTaskRepository()
	taskService := taskservices.NewTaskService(configs.DB, taskRepo, tagRepository, validate)
	myTaskControllers := taskControllers.NewTaskController(taskService)

	r := gin.Default()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)
	r.GET("/dashboard", middleware.AuthMiddleware, Controller.GetProfile)
	r.POST("/logout", Controller.Logout)
	r.POST("/task", middleware.AuthMiddleware, myTaskControllers.Create)
	r.PUT("/task/:id", middleware.AuthMiddleware, myTaskControllers.Update)
	r.DELETE("/task/:id", middleware.AuthMiddleware, myTaskControllers.Delete)

	err := r.Run()
	if err != nil {
		return
	}
}
