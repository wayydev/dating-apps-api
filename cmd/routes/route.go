package routes

import (
	"dating-apps/api/cmd/controllers"
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func Init(service *services.Service) {
	router := gin.Default()

	router.Use(utilities.ErrorHandle())
	router.NoRoute(utilities.ErrorNotFound())

	AuthRoute(router, controllers.NewAuthController(service))
	SwapRoute(router, controllers.NewSwapController(service), service)
	ProfileRoute(router, controllers.NewProfileController(service))

	router.Run(":8000")
}
