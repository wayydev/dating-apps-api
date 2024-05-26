package middlewares

import (
	"dating-apps/api/cmd/services"
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func LimiterMiddleware(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, _ := c.Get("auth")
		if err := service.SwapService.Limiter(auth.(*utilities.JWT).ID); err != nil {
			utilities.ErrorPanic(err)
			return
		}

		c.Next()
	}
}
