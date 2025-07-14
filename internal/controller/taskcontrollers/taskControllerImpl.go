package taskcontrollers

import (
	"Yattask/internal/common"
	"Yattask/internal/dto"
	"Yattask/internal/entities"
	"Yattask/internal/service/taskservices"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"net/http"
	"os"
	"strconv"
)

type TaskControllerImpl struct {
	service taskservices.TaskService
}

func NewTaskController(service taskservices.TaskService) TaskController {
	return &TaskControllerImpl{
		service: service,
	}
}

func (t *TaskControllerImpl) Create(c *gin.Context) {
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

func (t *TaskControllerImpl) Update(c *gin.Context) {
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

func (t *TaskControllerImpl) Delete(c *gin.Context) {
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

func (t *TaskControllerImpl) CreateWithAI(c *gin.Context) {
	var userReq dto.TaskWebCreateRequestWithAI
	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
		return
	}
	user, _ := c.Get("user")
	userId := user.(entities.User).ID
	systemMsg := fmt.Sprintf(`
Ubah kalimat user menjadi JSON valid tanpa komentar atau teks tambahan.

Format:
{
  "user_id": %d,
  "title": "string",
  "deadline": "string",
  "description": "string",
  "status": false,
  "tags": ["string"]
}

- Gunakan user_id = %d
- Gunakan double quote
- Jika status tidak disebut, default false
- Jika tidak ada tags, isi dengan []

Contoh:
"Bikin tugas matematika deadline besok, tag sekolah"
`, userId, userId)
	userMsg := userReq.Message
	model := "openrouter/cypher-alpha:free"

	client := openai.NewClient(option.WithBaseURL("https://openrouter.ai/api/v1"), option.WithAPIKey(os.Getenv("API_KEY")))
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMsg),
			openai.UserMessage(userMsg),
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to create task"})
		return
	}
	if len(chatCompletion.Choices) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "AI gave no response"})
		return
	}
	var message dto.TaskCreateRequest
	content := chatCompletion.Choices[0].Message.Content
	err = json.Unmarshal([]byte(content), &message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parsed"})
		return
	}

	taskResp, err := t.service.Create(c.Request.Context(), message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to create"})
	}
	response := common.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "task created",
		Message: taskResp,
	}
	c.JSON(http.StatusCreated, response)

}

func (t *TaskControllerImpl) GetByIDWithTask(c *gin.Context) {
	taskIdStr := c.Param("id")
	taskIdU64, _ := strconv.ParseUint(taskIdStr, 10, 64)
	taskId := uint(taskIdU64)
	user, _ := c.Get("user")
	userId := user.(entities.User).ID
	task, err := t.service.GetByIdWithTags(c.Request.Context(), taskId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to get"})
	}
	response := common.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "get task",
		Message: task,
	}
	c.JSON(http.StatusOK, response)
}
