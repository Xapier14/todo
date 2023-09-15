package auth

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xapier14/todo/internal/responses/check"
	"github.com/xapier14/todo/internal/utils/jwt"
)

// GetCheck godoc

// @Router /auth/check [get]
func GetCheck(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		// no access token cookie
		fmt.Println("No access token cookie")
		fmt.Println(err)
		c.AbortWithStatusJSON(401, check.GenerateNotLoggedInResponse("no_access_token_cookie"))
		return
	}
	jwtData, err := jwt.ParseJWT(accessToken)
	if err != nil {
		// invalid access token cookie
		fmt.Println("Invalid access token cookie")
		fmt.Println(err)
		c.AbortWithStatusJSON(401, check.GenerateNotLoggedInResponse("invalid_access_token_cookie"))
		return
	}

	c.JSON(200, check.GenerateLoggedInResponse(strconv.FormatFloat(jwtData["token_id"].(float64), 'f', -1, 64)))
}