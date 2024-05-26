package middlewares

import (
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func LimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			utilities.Error("Authorization", 401, nil, nil)
			return
		}

		authToken := authorization[len("Bearer "):]

		auth, err := utilities.ParseTokenJWT(authToken)

		if err != nil {
			utilities.ErrorPanic(err)
			return
		}

		c.Set("auth", auth)
		c.Next()
	}
}
