package credentials

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func GetUserCredentialByEmail(email string, userCredential *models.UserCredential) error {
	gormDB := db.GetDB()
	result := gormDB.First(&userCredential, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserCredentialByID(id uint, userCredential *models.UserCredential) error {
	gormDB := db.GetDB()
	result := gormDB.First(&userCredential, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ExistsUserCredentialByEmail(email string) bool {
	gormDB := db.GetDB()
	userCredential := models.UserCredential{}
	result := gormDB.First(&userCredential, "email = ?", email)
	return result.Error == nil
}