package main

import "C"
import (
	"Yattask/configs"
	"Yattask/internal/controller/taskControllers"
	"Yattask/internal/controller/userControllers"
	"Yattask/internal/repository/tagrepositories"
	"Yattask/internal/repository/taskrepositories"
	"Yattask/internal/repository/userrepositories"
	"Yattask/internal/router"
	"Yattask/internal/service/taskservices"
	"Yattask/internal/service/userservices"
	"github.com/go-playground/validator/v10"
)

func init() {
	configs.LoadEnv()
	configs.GetConnDB()
	configs.SyncTables(configs.DB)
}

func main() {
	validate := validator.New()
	MyUserRepo := userrepositories.NewUserRepository()
	MyUserService := userservices.NewUserService(configs.DB, MyUserRepo, validate)
	MyUserController := userControllers.NewUserController(MyUserService)
	MyTagRepo := tagrepositories.NewTagRepository()
	MyTaskRepo := taskrepositories.NewTaskRepository()
	MyTaskService := taskservices.NewTaskService(configs.DB, MyTaskRepo, MyTagRepo, validate)
	MyTaskControllers := taskControllers.NewTaskController(MyTaskService)

	r := router.AllRoutes(MyUserController, MyTaskControllers)

	err := r.Run()
	if err != nil {
		return
	}
}
