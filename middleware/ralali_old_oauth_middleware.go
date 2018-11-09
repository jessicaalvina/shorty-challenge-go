package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (m *DefaultMiddleware) RalaliOldOAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken := c.GetHeader("x-access-token")
		grantToken := c.GetHeader("x-grant-code")

		// user credential
		if "" != accessToken {
			fmt.Println(accessToken)
		}

		// client credential
		if "" != grantToken {
			fmt.Println(grantToken)
		}

		c.Next()

	}
}
