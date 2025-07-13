package app

import (
	"Yattask/configs"
	"Yattask/internal/di"
	"Yattask/internal/router"
	"log"
)

func RunServer() {
	userCtrl := di.InitializeUserControllers(configs.DB, configs.Validate)
	taskCtrl := di.InitializeTaskControllers(configs.DB, configs.Validate)

	r := router.AllRoutes(userCtrl, taskCtrl)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
