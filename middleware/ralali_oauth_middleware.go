package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ralali.com/services"
)

func (m *DefaultMiddleware) RalaliOAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		redisService := services.RedisService{}
		redisService, _ = redisService.Initialize()

		authorizationBearer := c.GetHeader("Authorization")
		accessToken := c.GetHeader("x-access-token")
		grantToken := c.GetHeader("x-grant-code")

		fmt.Println(authorizationBearer)
		fmt.Println(accessToken)
		fmt.Println(grantToken)

		if "" == accessToken && "" == grantToken {
			return
		}

		redisGrantKey := "oauthGrant:" + grantToken
		redisAccessKey := "oauthAccess:" + accessToken

		redisGrantData := redisService.Client.Get(redisGrantKey)
		redisAccessData := redisService.Client.Get(redisAccessKey)

		fmt.Println(redisGrantData.String())
		fmt.Println(redisAccessData.String())


		c.Next()

	}

}
