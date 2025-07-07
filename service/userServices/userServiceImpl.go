package userServices

import (
	"Yattask/dto/userDTO"
	"Yattask/helper"
	"Yattask/model"
	"Yattask/repository/userRepositories"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type UserServiceImpl struct {
	DB             *gorm.DB
	UserRepository userRepositories.UserRepository
	Validator      *validator.Validate
}

func NewUserService(db *gorm.DB, userRepository userRepositories.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		DB:             db,
		UserRepository: userRepository,
		Validator:      validate,
	}
}

func (u UserServiceImpl) Register(ctx context.Context, user userDTO.UserRegisterRequest) (userDTO.UserServiceResponse, error) {
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

func (u UserServiceImpl) Login(ctx context.Context, user userDTO.UserLoginRequest) (string, error) {
	err := u.Validator.Struct(user)
	if err != nil {
		return "", errors.New("invalid login request")
	}
	const tokenExpired = time.Hour * 24
	var tokenString string
	errTx := u.DB.Transaction(func(tx *gorm.DB) error {
		getedUser, err := u.UserRepository.GetByUsername(ctx, tx, user.Username)
		if err != nil {
			return errors.New("user not found")
		}
		err = bcrypt.CompareHashAndPassword([]byte(getedUser.Password), []byte(user.Password))
		if err != nil {
			return errors.New("password is incorrect")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": getedUser.ID,
			"exp": time.Now().Add(tokenExpired).Unix(),
		})

		generatedStringToken, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if errToken != nil {
			return errors.New("token is invalid")
		}

		tokenString = generatedStringToken
		return nil

	})
	if errTx != nil {
		return "", errors.New("failed to login")
	}
	return tokenString, nil
}

func (u UserServiceImpl) GetProfileWithTaskAndTag(ctx context.Context, userID uint) (userDTO.UserServiceResponseGETWithTaskTag, error) {
	var user model.User
	errTx := u.DB.Transaction(func(tx *gorm.DB) error {
		getedUser, err := u.UserRepository.GetByIdWithTaskAndTag(ctx, tx, userID)
		if err != nil {
			return err
		}
		user = getedUser
		return nil

	})
	if errTx != nil {
		return userDTO.UserServiceResponseGETWithTaskTag{}, errTx

	}

	return helper.ToUserServiceResponseWithTaskTag(user), nil
}

func (u UserServiceImpl) GetByID(ctx context.Context, userID uint) (userDTO.UserServiceResponseGETWithTaskTag, error) {
	var user model.User
	errTX := u.DB.Transaction(func(tx *gorm.DB) error {
		getedUser, err := u.UserRepository.GetById(ctx, tx, userID)
		if err != nil {
			return err
		}
		user = getedUser
		return nil
	})
	if errTX != nil {
		return userDTO.UserServiceResponseGETWithTaskTag{}, errTX
	}
	return helper.ToUserServiceResponseWithTaskTag(user), nil
}
