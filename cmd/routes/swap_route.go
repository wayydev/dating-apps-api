package routes

import (
	"dating-apps/api/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func SwapRoute(router *gin.Engine, controller *controllers.SwapController) {
	routes := router.Group("/swap")
	{
		routes.GET("/", controller.Find)
		routes.PUT("/like/:swap_user_id", controller.Like)
		routes.PUT("/pass/:swap_user_id", controller.Pass)
	}
}
