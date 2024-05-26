package controllers

import (
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SwapController struct {
	Service *services.Service
}

func NewSwapController(service *services.Service) *SwapController {
	return &SwapController{service}
}

func (s SwapController) Find(c *gin.Context) {
	auth, _ := c.Get("auth")
	user, err := s.Service.SwapService.Find(auth.(*utilities.JWT).ID)

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Success get nearby peoples", 200, user)
}

func (s SwapController) Like(c *gin.Context) {
	auth, _ := c.Get("auth")
	swapUserID, _ := strconv.Atoi(c.Param("swap_id"))

	user, err := s.Service.SwapService.Like(auth.(*utilities.JWT), uint(swapUserID))

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Success get nearby peoples", 200, user)
}

func (s SwapController) Pass(c *gin.Context) {
	auth, _ := c.Get("auth")
	swapUserID, _ := strconv.Atoi(c.Param("swap_id"))

	user, err := s.Service.SwapService.Pass(auth.(*utilities.JWT).ID, uint(swapUserID))

	if err != nil {
		utilities.ErrorPanic(err)
		return
	}

	utilities.ResponseJSON(c, "Success get nearby peoples", 200, user)
}
