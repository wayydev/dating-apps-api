package routes

import (
	"dating-apps/api/cmd/controllers"
	"dating-apps/api/cmd/middlewares"
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func Init(service *services.Service) {
	router := gin.Default()
	router.Use(utilities.ErrorHandle())

	AuthRoute(router, controllers.NewAuthController(service))

	router.Use(middlewares.AuthMiddleware())
	{
		router.Use(middlewares.LimiterMiddleware(service))
		{
			SwapRoute(router, controllers.NewSwapController(service))
		}

		ProfileRoute(router, controllers.NewProfileController(service))
	}

	router.NoRoute(utilities.ErrorNotFound)
	router.Run(":8000")
}
