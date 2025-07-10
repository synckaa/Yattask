package helper

import (
	"Yattask/dto/taskDTO"
	"Yattask/dto/userDTO"
	"Yattask/model"
)

func ToUserServiceResponse(user model.User) userDTO.UserServiceResponse {
	return userDTO.UserServiceResponse{
		Model:    user.Model,
		Username: user.Username,
	}

}

func ToUserServiceResponseWithTaskTag(user model.User) userDTO.UserServiceResponseGETWithTaskTag {
	var tags []string
	for _, tag := range user.Tags {
		tags = append(tags, tag.Name)
	}

	var tasks []taskDTO.TaskWithTagsResponse
	for _, task := range user.Tasks {
		var tagName []string
		for _, tag := range task.Tags {
			tagName = append(tagName, tag.Name)
		}
		tasks = append(tasks, taskDTO.TaskWithTagsResponse{
			Model:       task.Model,
			Title:       task.Title,
			Deadline:    task.Deadline,
			Description: task.Description,
			Status:      task.Status,
			Tags:        tagName,
		})
	}
	return userDTO.UserServiceResponseGETWithTaskTag{
		Model:    user.Model,
		Username: user.Username,
		Tasks:    tasks,
		Tags:     tags,
	}
}
