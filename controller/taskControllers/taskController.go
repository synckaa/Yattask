package taskControllers

import "github.com/gin-gonic/gin"

type TaskController interface {
	Create(c *gin.Context)
}
