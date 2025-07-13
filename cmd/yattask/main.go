package main

import (
	"Yattask/internal/app"
	"Yattask/internal/bootstrap"
)

func main() {
	bootstrap.InitApp()
	app.RunServer()
}
