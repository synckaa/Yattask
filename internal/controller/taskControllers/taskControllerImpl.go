package taskControllers

import (
	"Yattask/internal/common"
	"Yattask/internal/dto"
	"Yattask/internal/entities"
	"Yattask/internal/service/taskservices"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskControlerImpl struct {
	service taskservices.TaskService
}

func NewTaskController(service taskservices.TaskService) TaskController {
	return TaskControlerImpl{
		service: service,
	}
}

func (t TaskControlerImpl) Create(c *gin.Context) {
	var task dto.TaskWebCreateRequest
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
	}
	user, _ := c.Get("user")
	userId := user.(entities.User).ID
	taskReq := dto.TaskCreateRequest{
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

	response := common.WebStandardResponse{
		Code:    http.StatusCreated,
		Status:  "task created",
		Message: taskResp,
	}
	c.JSON(http.StatusCreated, response)

}

func (t TaskControlerImpl) Update(c *gin.Context) {
	var task dto.TaskWebUpdateRequest
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
	}
	user, _ := c.Get("user")
	userId := user.(entities.User).ID
	taskIdStr := c.Param("id")
	taskIdU64, _ := strconv.ParseUint(taskIdStr, 10, 64)
	taskId := uint(taskIdU64)
	taskReq := dto.TaskUpdateRequest{
		ID:          taskId,
		UserID:      userId,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
		Tags:        task.Tags,
	}
	taskResp, err := t.service.Update(c.Request.Context(), taskReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to update"})
	}
	response := common.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "task updated",
		Message: taskResp,
	}
	c.JSON(http.StatusOK, response)
}

func (t TaskControlerImpl) Delete(c *gin.Context) {
	user, _ := c.Get("user")
	userId := user.(entities.User).ID
	taskIdStr := c.Param("id")
	taskIdU64, _ := strconv.ParseUint(taskIdStr, 10, 64)
	taskId := uint(taskIdU64)
	err := t.service.Delete(c.Request.Context(), taskId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to delete"})
		return
	}
	response := common.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "successfully deleted",
	}
	c.JSON(http.StatusOK, response)
}
