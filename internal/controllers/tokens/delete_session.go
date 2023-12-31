package tokens

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func DeleteSessionTokenById(id uint) error {
	gormDB := db.GetDB()

	var session models.UserSession
	result := gormDB.Where(id).Delete(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteSessionTokenByToken(token string) error {
	gormDB := db.GetDB()

	var session models.UserSession
	result := gormDB.Where("token = ?", token).Delete(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}