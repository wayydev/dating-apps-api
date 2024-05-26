package routes

import (
	"dating-apps/api/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func SwapRoute(router *gin.Engine, controller *controllers.SwapController) {
	routes := router.Group("/")
	{
		routes.GET("/swap", controller.Find)
		routes.PUT("/like", controller.Like)
		routes.PUT("/pass", controller.Pass)
	}
}
