package routes

import (
	"dating-apps/api/cmd/controllers"
	"dating-apps/api/pkg/injectors"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func Init(service *injectors.Service) {
	router := gin.Default()
	router.Use(utilities.ErrorHandle())

	AuthRoute(router, controllers.NewAuthController(service))

	router.NoRoute(utilities.ErrorNotFound)
	router.Run(":3000")
}
