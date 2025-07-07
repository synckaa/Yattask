package helper

import (
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
	return userDTO.UserServiceResponseGETWithTaskTag{
		Model:    user.Model,
		Username: user.Username,
		Tasks:    user.Tasks,
		Tags:     user.Tags,
	}
}
