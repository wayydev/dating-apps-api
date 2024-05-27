package middlewares

import (
	"dating-apps/api/pkg/utilities"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			errResponse := utilities.Error("Authorization header missing", 401, nil, nil)
			c.JSON(errResponse.Status, errResponse)
			c.Abort()
			return
		}

		authToken := authorization[len("Bearer "):]

		auth, err := utilities.ParseTokenJWT(authToken)

		if err != nil || auth == nil {
			errResponse := utilities.Error("Invalid token", 401, nil, err)
			c.JSON(errResponse.Status, errResponse)
			c.Abort()
			return
		}

		c.Set("auth", auth)
		c.Next()
	}
}
