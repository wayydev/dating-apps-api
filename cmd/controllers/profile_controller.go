package controllers

import (
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	Service *services.Service
}

func NewProfileController(service *services.Service) *ProfileController {
	return &ProfileController{service}
}

func (s ProfileController) Get(c *gin.Context) {
	auth, _ := c.Get("auth")
	user, err := s.Service.ProfileService.GetProfile(auth.(*utilities.JWT).ID)

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Success get Profile", 200, user)
}

func (s ProfileController) Notification(c *gin.Context) {
	auth, _ := c.Get("auth")
	user, err := s.Service.ProfileService.GetNotification(auth.(*utilities.JWT).ID)

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Success get notification", 200, user)
}
