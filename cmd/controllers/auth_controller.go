package controllers

import (
	"dating-apps/api/cmd/requests"
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.Service
}

func NewAuthController(service *services.Service) *AuthController {
	return &AuthController{service}
}

func (s AuthController) Login(c *gin.Context) {
	var req *requests.Login
	c.ShouldBindJSON(&req)

	user, err := s.Service.AuthService.Login(req)

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Login Success", 200, user)
}

func (s AuthController) Registration(c *gin.Context) {
	req := requests.Registration{}
	c.ShouldBindJSON(&req)

	err := s.Service.AuthService.Registration(&req)

	if err != nil {
		utilities.ErrorPanic(err)
	}

	utilities.ResponseJSON(c, "Registration Success", 201, nil)
}
