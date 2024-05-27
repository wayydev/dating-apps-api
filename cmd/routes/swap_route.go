package routes

import (
	"dating-apps/api/cmd/controllers"
	"dating-apps/api/cmd/middlewares"
	"dating-apps/api/cmd/services"

	"github.com/gin-gonic/gin"
)

func SwapRoute(router *gin.Engine, controller *controllers.SwapController, service *services.Service) {
	routes := router.Group("/swap")
	routes.Use(middlewares.AuthMiddleware())
	routes.Use(middlewares.LimiterMiddleware(service))
	{
		routes.GET("/", controller.Find)
		routes.PUT("/like/:swap_user_id", controller.Like)
		routes.PUT("/pass/:swap_user_id", controller.Pass)
	}
}
