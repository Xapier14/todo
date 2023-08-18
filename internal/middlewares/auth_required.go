package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xapier14/todo/internal/utils/jwt"
)

func RequireAuth(returnJson any) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			// no access token cookie
			fmt.Println("No access token cookie")
			fmt.Println(err)
			c.AbortWithStatusJSON(401, returnJson)
			return
		}
		jwtData, err := jwt.ParseJWT(accessToken)
		if err != nil {
			// invalid access token cookie
			fmt.Println("Invalid access token cookie")
			fmt.Println(err)
			c.AbortWithStatusJSON(401, returnJson)
			return
		}
		c.Set("access_token_data", jwtData)
		c.Set("token_id", jwtData["token_id"])
		c.Set("user_id", jwtData["user_id"])
		c.Set("user_email", jwtData["email"])
		c.Next()
	}
}