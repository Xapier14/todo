package credentials

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func DeleteUserCredential(userCredential *models.UserCredential) error {
	gormDB := db.GetDB()
	result := gormDB.Delete(&userCredential)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserCredentialByEmail(email string) error {
	gormDB := db.GetDB()
	result := gormDB.Delete(&models.UserCredential{}, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}