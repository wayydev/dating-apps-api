package routes

import (
	"dating-apps/api/cmd/controllers"
	"dating-apps/api/cmd/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine, controller *controllers.ProfileController) {
	routes := router.Group("/profile")
	routes.Use(middlewares.AuthMiddleware())
	{
		routes.GET("/", controller.Get)
		routes.GET("/notifications", controller.Notification)
	}
}
