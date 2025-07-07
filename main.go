package main

import (
	"Yattask/config"
	"Yattask/controller/userControllers"
	"Yattask/middleware"
	"Yattask/repository/userRepositories"
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

	r := gin.Default()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)
	r.GET("/dashboard", middleware.AuthMiddleware, Controller.GetProfile)
	r.POST("/logout", Controller.Logout)

	err := r.Run()
	if err != nil {
		return
	}
}
