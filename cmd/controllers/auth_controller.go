package controllers

import (
	"dating-apps/api/cmd/requests"
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/injectors"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(service *injectors.Service) *AuthController {
	return &AuthController{
		AuthService: service.AuthService,
	}
}

func (s AuthController) Login(c *gin.Context) {
	var req *requests.Login
	c.ShouldBindJSON(&req)

	user, err := s.AuthService.Login(req)

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Login Success", 200, user)
}

func (s AuthController) Registration(c *gin.Context) {
	var req *requests.Registration
	err := s.AuthService.Registration(req)

	if err != nil {
		utilities.ErrorPanic(err)
	}

	utilities.ResponseJSON(c, "Registration Success", 201, nil)
}
