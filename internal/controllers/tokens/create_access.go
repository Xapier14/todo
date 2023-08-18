package tokens

import (
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/utils/jwt"
)

func CreateAccessToken(userCredential *models.UserCredential, session *models.UserSession) (string, error) {
	mapClaims := map[string]interface{}{
		"email": userCredential.Email,
		"user_id": userCredential.ID,
		"token_id": session.ID,
	}
	jwt, err := jwt.GenerateDefaultExpiringJWT(mapClaims)
	if err != nil {
		return "", err
	}
	return jwt, nil
}