package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/tokens"
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses/general"
	"github.com/xapier14/todo/internal/responses/login"
	"github.com/xapier14/todo/internal/utils/hashing"
)

// PostLogin godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json

// @Router /auth/login [post]
func PostLogin(c *gin.Context) {
	gormDB := db.GetDB()
	email := c.PostForm("email")
	password := c.PostForm("password")

	salt := hashing.GenerateSalt()
	userCredential := models.UserCredential{}
	result := gormDB.First(&userCredential, "email = ?", email)
	if result.Error == nil {
		salt = userCredential.Salt
	}
	hashedPassword := hashing.HashPassword(password, salt)

	if (result.Error != nil) || (hashedPassword != userCredential.PasswordHash) {
		c.JSON(401, login.GenerateInvalidCredentialsResponse())
		return
	}

	sessionToken, err := tokens.CreateSessionToken(userCredential.ID)
	if err != nil {
		c.JSON(500, general.GenerateDetailedInternalServerErrorResponse("Could not issue session token"))
		return
	}

	accessToken, err := tokens.CreateAccessToken(&userCredential)
	if err != nil {
		c.JSON(500, general.GenerateDetailedInternalServerErrorResponse("Could not issue access token"))
		return
	}

	c.SetCookie("access_token", accessToken, 60*60*24*30, "/", "", false, true)

	c.JSON(200, login.GenerateSuccessfulLoginResponse(sessionToken.Token))
}