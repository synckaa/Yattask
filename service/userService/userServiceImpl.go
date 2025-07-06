package userService

import (
	"Yattask/dto/userDTO"
	"Yattask/helper"
	"Yattask/model"
	"Yattask/repository/userRepository"
	"context"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userServiceImpl struct {
	DB             *gorm.DB
	UserRepository userRepository.UserRepository
	Validator      *validator.Validate
}

func (u userServiceImpl) Register(ctx context.Context, user userDTO.UserRegisterRequest) (userDTO.UserServiceResponse, error) {
	err := u.Validator.Struct(user)
	if err != nil {
		return userDTO.UserServiceResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDTO.UserServiceResponse{}, err
	}
	userReq := model.User{
		Username: user.Username,
		Password: string(hash),
	}
	var userResp model.User
	errTx := u.DB.Transaction(func(tx *gorm.DB) error {
		createdUser, err := u.UserRepository.Create(ctx, tx, userReq)
		if err != nil {
			return err
		}
		userResp = createdUser
		return nil

	})
	if errTx != nil {
		return userDTO.UserServiceResponse{}, errTx
	}
	return helper.ToUserServiceResponse(userResp), nil

}

func (u userServiceImpl) Login(ctx context.Context, user userDTO.UserLoginRequest) (userDTO.UserServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) UpdateProfile(ctx context.Context, user userDTO.UserLoginRequest) (userDTO.UserServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) GetProfileWithTaskAndTag(ctx context.Context, userID int) (userDTO.UserServiceResponseGETWithTaskTag, error) {
	//TODO implement me
	panic("implement me")
}
