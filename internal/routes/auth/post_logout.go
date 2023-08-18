package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/tokens"
	"github.com/xapier14/todo/internal/responses/logout"
)

func PostLogout(c *gin.Context) {
	tokenId := c.GetUint("token_id")
	err := tokens.DeleteSessionTokenById(tokenId)
	if err != nil {
		c.AbortWithStatusJSON(500, logout.GenerateLogoutErrorResponse())
		return
	}

	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.JSON(200, logout.GenerateSuccessfulLoginResponse())
}