package auth

import (
	"boilerplate/configs"
	"boilerplate/models"
	"boilerplate/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type service struct {
	repository Repository
	entity     Entity
}

func NewServiceRegister() *service {
	_repository := NewRepository(configs.GetDB(), configs.GetRedis())
	return &service{repository: _repository}
}

func (s *service) Register(c *gin.Context) {
	// Validate input
	var input InputRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, utils.ValidatorError(err))
		return
	}

	users := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := s.repository.Register(&users); err != nil {
		utils.APIResponse(c, err.Error(), http.StatusConflict, nil)
		return
	}

	utils.SendMail(users.Fullname, users.Email, "User Activation", "template_register", "123")

	utils.APIResponse(c, "Register new account successfully", http.StatusCreated, nil)
}

func (s *service) Login(c *gin.Context) {
	var input InputLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidatorError(err))
		return
	}

	users := models.User{
		Email:    strings.ToLower(input.Email),
		Password: input.Password,
	}

	result, err := s.repository.Login(&users)
	if err != nil {
		utils.APIResponse(c, err.Error(), http.StatusUnauthorized, nil)
		return
	}

	accessUUID := uuid.New().String()
	accessExpired := time.Unix(time.Now().Add(time.Minute*15).Unix(), 0)
	accessTokenData := map[string]interface{}{"access_uuid": accessUUID}
	accessToken, err := utils.Sign(accessTokenData, true, "ACCESS_SECRET", accessExpired.Unix())

	if err := s.repository.SaveToken(accessUUID, result.ID, accessExpired.Sub(time.Now())); err != nil {
		utils.APIResponse(c, err.Error(), http.StatusUnauthorized, nil)
		return
	}

	refreshUUID := uuid.New().String()
	refreshExpired := time.Unix(time.Now().Add(time.Hour*24).Unix(), 0)
	refreshTokenData := map[string]interface{}{"refresh_uuid": refreshUUID}
	refreshToken, err := utils.Sign(refreshTokenData, true, "REFERSH_SECRET", refreshExpired.Unix())

	if err := s.repository.SaveToken(refreshUUID, result.ID, refreshExpired.Sub(time.Now())); err != nil {
		utils.APIResponse(c, err.Error(), http.StatusUnauthorized, nil)
		return
	}

	response := map[string]interface{}{
		"user": UserResponse{
			Fullname: result.Fullname,
			Email:    result.Email,
		},
		"token": TokenResponse{
			AccessToken:    accessToken,
			AccessExpired:  accessExpired,
			RefreshToken:   refreshToken,
			RefreshExpired: refreshExpired,
		},
	}
	fmt.Println(response)
	utils.APIResponse(c, "Login successfully", http.StatusOK, response)

}
