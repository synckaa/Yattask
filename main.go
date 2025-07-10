package main

import (
	"Yattask/config"
	"Yattask/controller/taskControllers"
	"Yattask/controller/userControllers"
	"Yattask/middleware"
	"Yattask/repository/tagRepositories"
	"Yattask/repository/taskRepositories"
	"Yattask/repository/userRepositories"
	"Yattask/service/taskServices"
	"Yattask/service/userServices"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	config.LoadEnv()
	config.GetConnDB()
	config.SyncTables(config.DB)
}

func main() {
	validate := validator.New()
	repo := userRepositories.NewUserRepository()
	Service := userServices.NewUserService(config.DB, repo, validate)
	Controller := userControllers.NewUserController(Service)
	tagRepository := tagRepositories.NewTagRepository()
	taskRepo := taskRepositories.NewTaskRepository()
	taskService := taskServices.NewTaskService(config.DB, taskRepo, tagRepository, validate)
	myTaskControllers := taskControllers.NewTaskController(taskService)

	r := gin.Default()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)
	r.GET("/dashboard", middleware.AuthMiddleware, Controller.GetProfile)
	r.POST("/logout", Controller.Logout)
	r.POST("/task", middleware.AuthMiddleware, myTaskControllers.Create)

	err := r.Run()
	if err != nil {
		return
	}
}
