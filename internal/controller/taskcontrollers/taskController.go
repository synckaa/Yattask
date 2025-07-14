package taskcontrollers

import "github.com/gin-gonic/gin"

type TaskController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	CreateWithAI(c *gin.Context)
	GetByIDWithTask(c *gin.Context)
}
