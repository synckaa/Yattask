package taskControllers

import (
	"Yattask/dto"
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
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	taskReq := taskDTO.TaskCreateUpdateRequest{
		UserID:      userId,
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

	response := dto.WebStandardResponse{
		Code:    http.StatusCreated,
		Status:  "task created",
		Message: taskResp,
	}
	c.JSON(http.StatusCreated, response)

}
