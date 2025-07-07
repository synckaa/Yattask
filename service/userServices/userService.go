package userServices

import (
	"Yattask/dto/userDTO"
	"context"
)

type UserService interface {
	Register(ctx context.Context, user userDTO.UserRegisterRequest) (userDTO.UserServiceResponse, error)
	Login(ctx context.Context, user userDTO.UserLoginRequest) (string, error)
	GetProfileWithTaskAndTag(ctx context.Context, userID uint) (userDTO.UserServiceResponseGETWithTaskTag, error)
	GetByID(ctx context.Context, userID uint) (userDTO.UserServiceResponseGETWithTaskTag, error)
}
