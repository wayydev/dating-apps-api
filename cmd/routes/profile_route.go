package routes

import (
	"dating-apps/api/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine, controller *controllers.ProfileController) {
	routes := router.Group("/profile")
	{
		routes.GET("/", controller.Get)
	}
}
