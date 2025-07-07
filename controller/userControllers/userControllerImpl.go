package userControllers

import (
	"Yattask/dto"
	"Yattask/dto/userDTO"
	"Yattask/model"
	"Yattask/service/userServices"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	Service userServices.UserService
}

func NewUserController(service userServices.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (u UserControllerImpl) Register(c *gin.Context) {
	var userReq userDTO.UserRegisterRequest
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
	}
	registered, err := u.Service.Register(c.Request.Context(), userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to register user"})
	}
	response := dto.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: registered,
	}
	c.JSON(http.StatusOK, response)

}

func (u UserControllerImpl) Login(c *gin.Context) {
	tokenValidate, _ := c.Cookie("token")
	if tokenValidate != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "You are Logged In"})
		return
	}
	var userReq userDTO.UserLoginRequest
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to binding"})
		return
	}
	token, err := u.Service.Login(c.Request.Context(), userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to login"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 3600*24, "/", "", false, true)
	response := dto.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "Login successful",
	}
	c.JSON(http.StatusOK, response)
}

func (u UserControllerImpl) GetProfile(c *gin.Context) {
	loginedUser, _ := c.MustGet("user").(model.User)
	user, err := u.Service.GetProfileWithTaskAndTag(c.Request.Context(), loginedUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to Get Task"})
	}
	response := dto.WebStandardResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: user,
	}
	c.JSON(http.StatusOK, response)
}

func (u UserControllerImpl) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})

}
