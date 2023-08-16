package tokens

import (
	"time"

	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/utils/hashing"
)

const (
	// DefaultSessionTokenExpiry is the default expiry time for a session token in hours
	DefaultSessionTokenExpiry = 24 * 7
)

func CreateSessionToken(userID uint) (*models.UserSession, error) {
	gormDB := db.GetDB()

	sessionToken := hashing.GenerateUUID()
	expiresAt := time.Now().Add(DefaultSessionTokenExpiry * time.Hour)

	session := models.UserSession{
		UserID:		   userID,
		Token:		   sessionToken,
		ExpiresAt:	   expiresAt,
	}

	result := gormDB.Create(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}

func CreateSessionTokenFromUserSession(userSession *models.UserSession) (*models.UserSession, error) {
	gormDB := db.GetDB()

	sessionToken := hashing.GenerateUUID()
	expiresAt := time.Now().Add(DefaultSessionTokenExpiry * time.Hour)

	session := models.UserSession{
		UserID:		   	userSession.UserID,
		Token:		   	sessionToken,
		ExpiresAt:		expiresAt,
		SessionInfo:  	userSession.SessionInfo,
	}

	result := gormDB.Create(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}