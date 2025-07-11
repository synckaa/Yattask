package userservices

import (
	"Yattask/internal/dto"
	"context"
)

type UserService interface {
	Register(ctx context.Context, user dto.UserRegisterRequest) (dto.UserServiceResponse, error)
	Login(ctx context.Context, user dto.UserLoginRequest) (string, error)
	GetProfileWithTaskAndTag(ctx context.Context, userID uint) (dto.UserServiceResponseGETWithTaskTag, error)
}
