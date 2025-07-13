package main

import "C"
import (
	"Yattask/configs"
	"Yattask/internal/di"
	"Yattask/internal/router"
)

func init() {
	configs.LoadEnv()
	configs.GetConnDB()
	configs.SyncTables(configs.DB)
	configs.NewValidator()
}

func main() {
	MyUserController := di.InitializeUserControllers(configs.DB, configs.Validate)
	MyTaskControllers := di.InitializeTaskControllers(configs.DB, configs.Validate)

	r := router.AllRoutes(MyUserController, MyTaskControllers)

	err := r.Run()
	if err != nil {
		return
	}
}
