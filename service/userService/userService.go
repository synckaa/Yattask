package userService

import (
	"Yattask/dto/userDTO"
	"context"
)

type UserService interface {
	Register(ctx context.Context, user userDTO.UserRegisterRequest) (userDTO.UserServiceResponse, error)
	Login(ctx context.Context, user userDTO.UserLoginRequest) (userDTO.UserServiceResponse, error)
	UpdateProfile(ctx context.Context, user userDTO.UserLoginRequest) (userDTO.UserServiceResponse, error)
	GetProfileWithTaskAndTag(ctx context.Context, userID int) (userDTO.UserServiceResponseGETWithTaskTag, error)
}
