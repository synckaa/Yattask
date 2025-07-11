package helper

import (
	"Yattask/internal/dto"
	"Yattask/internal/entities"
)

func ToUserServiceResponse(user entities.User) dto.UserServiceResponse {
	return dto.UserServiceResponse{
		Model:    user.Model,
		Username: user.Username,
	}

}

func ToUserServiceResponseWithTaskTag(user entities.User) dto.UserServiceResponseGETWithTaskTag {
	var tags []string
	for _, tag := range user.Tags {
		tags = append(tags, tag.Name)
	}

	var tasks []dto.TaskWithTagsResponse
	for _, task := range user.Tasks {
		var tagName []string
		for _, tag := range task.Tags {
			tagName = append(tagName, tag.Name)
		}
		tasks = append(tasks, dto.TaskWithTagsResponse{
			Model:       task.Model,
			Title:       task.Title,
			Deadline:    task.Deadline,
			Description: task.Description,
			Status:      task.Status,
			Tags:        tagName,
		})
	}
	return dto.UserServiceResponseGETWithTaskTag{
		Model:    user.Model,
		Username: user.Username,
		Tasks:    tasks,
		Tags:     tags,
	}
}
