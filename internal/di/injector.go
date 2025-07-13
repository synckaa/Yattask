//go:build wireinject
// +build wireinject

package di

import (
	"Yattask/internal/controller/taskControllers"
	"Yattask/internal/controller/userControllers"
	"Yattask/internal/repository/tagrepositories"
	"Yattask/internal/repository/taskrepositories"
	"Yattask/internal/repository/userrepositories"
	"Yattask/internal/service/taskservices"
	"Yattask/internal/service/userservices"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserControllers(db *gorm.DB, validation *validator.Validate) userControllers.UserController {
	wire.Build(
		userrepositories.NewUserRepository,
		userservices.NewUserService,
		userControllers.NewUserController,
	)
	return nil
}

func InitializeTaskControllers(db *gorm.DB, validation *validator.Validate) taskControllers.TaskController {
	wire.Build(
		taskrepositories.NewTaskRepository,
		tagrepositories.NewTagRepository,
		taskservices.NewTaskService,
		taskControllers.NewTaskController)
	return nil
}
