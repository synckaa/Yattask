//go:build wireinject
// +build wireinject

package di

import (
	"Yattask/internal/controller/taskcontrollers"
	"Yattask/internal/controller/usercontrollers"
	"Yattask/internal/repository/tagrepositories"
	"Yattask/internal/repository/taskrepositories"
	"Yattask/internal/repository/userrepositories"
	"Yattask/internal/service/taskservices"
	"Yattask/internal/service/userservices"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserControllers(db *gorm.DB, validation *validator.Validate) usercontrollers.UserController {
	wire.Build(
		userrepositories.NewUserRepository,
		userservices.NewUserService,
		usercontrollers.NewUserController,
	)
	return nil
}

func InitializeTaskControllers(db *gorm.DB, validation *validator.Validate) taskcontrollers.TaskController {
	wire.Build(
		taskrepositories.NewTaskRepository,
		tagrepositories.NewTagRepository,
		taskservices.NewTaskService,
		taskcontrollers.NewTaskController)
	return nil
}
