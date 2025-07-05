package user

import (
	"Yattask/dto/user"
	"context"
)

type UserService interface {
	Register(ctx context.Context, user *user.UserRegisterRequest) (, error)
}
