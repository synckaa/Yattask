package helper

import (
	"Yattask/dto/userDTO"
	"Yattask/model"
)

func ToUserServiceResponse(user model.User) userDTO.UserServiceResponse {
	return userDTO.UserServiceResponse{
		Username: user.Username,
		Password: user.Password,
	}

}
