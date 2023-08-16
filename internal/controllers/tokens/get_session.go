package tokens

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)


func GetSessionTokenByToken(token string) (*models.UserSession, error) {
	gormDB := db.GetDB()

	var session models.UserSession
	result := gormDB.Where("token = ?", token).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}