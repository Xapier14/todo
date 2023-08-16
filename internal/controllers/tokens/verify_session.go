package tokens

import (
	"time"

	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func CheckIfSessionTokenExists(token string) bool {
	gormDB := db.GetDB()

	var session models.UserSession
	result := gormDB.Where("token = ?", token).First(&session)
	return result.Error == nil
}

func CheckIfSessionTokenIsValid(token string) bool {
	gormDB := db.GetDB()

	var session models.UserSession
	result := gormDB.Where("token = ?", token).First(&session)
	if result.Error != nil {
		return false
	}

	return session.ExpiresAt.After(time.Now())
}