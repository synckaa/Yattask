// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

package main

import (
	_ "Yattask/cmd/yattask/docs"
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

func main() {
	MyUserController := di.InitializeUserControllers(configs.DB, configs.Validate)
	MyTaskControllers := di.InitializeTaskControllers(configs.DB, configs.Validate)

	r := router.AllRoutes(MyUserController, MyTaskControllers)

	err := r.Run()
	if err != nil {
		return
	}
}
