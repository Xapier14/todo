package credentials

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func SaveUserCredential(userCredential *models.UserCredential) error {
	gormDB := db.GetDB()
	result := gormDB.Save(&userCredential)
	if result.Error != nil {
		return result.Error
	}
	return nil
}