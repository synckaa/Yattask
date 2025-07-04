package main

import (
	"Yattask/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.GetConnDB()
	config.SyncTables(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run()
}
