package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/credentials"
	"github.com/xapier14/todo/internal/controllers/tokens"
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses/login"
)

type refreshRequest struct {
	SessionToken string `json:"session_token"`
}

// PostRefresh godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Accept json
// @Produce json

// @Router /auth/refresh [post]
func PostRefresh(c *gin.Context) {
	var request refreshRequest
	c.BindJSON(&request)

	if request.SessionToken == "" {
		c.JSON(400, gin.H{
			"error": "session_token is required",
		})
		return
	}

	userSession, err := tokens.GetSessionTokenByToken(request.SessionToken)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid session token",
		})
		return
	}

	userCredential := models.UserCredential{}
	err = credentials.GetUserCredentialByID(userSession.UserID, &userCredential)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not find user",
		})
		return
	}

	err = tokens.DeleteSessionTokenByToken(request.SessionToken)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not renew session token",
		})
		return
	}

	sessionToken, err := tokens.CreateSessionTokenFromUserSession(userSession)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not issue session token",
		})
		return
	}

	accessToken, err := tokens.CreateAccessToken(&userCredential)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not issue access token",
		})
		return
	}

	c.SetCookie("access_token", accessToken, 60*60*24*30, "/", "", false, true)

	c.JSON(200, login.GenerateSuccessfulRefreshResponse(sessionToken.Token))
}