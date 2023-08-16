package auth

import (
	"github.com/gin-gonic/gin"
	users "github.com/xapier14/todo/internal/controllers/credentials"
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/utils/format"
	"github.com/xapier14/todo/internal/utils/hashing"
)

func PostSignup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !format.IsEmailValid(email) {
		c.JSON(400, gin.H{
			"message": "Invalid email",
		})
		return
	}

	if !format.IsPasswordValid(password) {
		c.JSON(400, gin.H{
			"message": "Invalid password",
		})
		return
	}

	salt := hashing.GenerateSalt()
	hashedPassword := hashing.HashPassword(password, salt)

	exists := users.ExistsUserCredentialByEmail(email)
	if exists {
		c.JSON(409, gin.H{
			"message": "Email already exists",
		})
		return
	}

	userCredential := models.UserCredential{
		Email:        email,
		PasswordHash: hashedPassword,
		Salt:         salt,
	}

	err := users.SaveUserCredential(&userCredential)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Signup successful",
	})
}