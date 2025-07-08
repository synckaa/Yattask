package taskControllers

import (
	"Yattask/dto/taskDTO"
	"Yattask/model"
	"Yattask/service/taskServices"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskControlerImpl struct {
	service taskServices.TaskService
}

func NewTaskController(service taskServices.TaskService) TaskController {
	return TaskControlerImpl{
		service: service,
	}
}

func (t TaskControlerImpl) Create(c *gin.Context) {
	var task taskDTO.TaskWebRequest
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
	}
	user := c.MustGet("user").(model.User)
	taskReq := taskDTO.TaskCreateUpdateRequest{
		UserID:      user.ID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
		Tags:        task.Tags,
	}
	taskResp, err := t.service.Create(c.Request.Context(), taskReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to create"})
	}
	c.JSON(http.StatusCreated, taskResp)

}
