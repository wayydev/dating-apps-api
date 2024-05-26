package routes

import (
	"dating-apps/api/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine, controller *controllers.AuthController) {
	routes := router.Group("/")
	{
		routes.POST("/login", controller.Login)
		routes.POST("/registration", controller.Registration)
	}
}
