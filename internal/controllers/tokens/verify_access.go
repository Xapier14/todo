package tokens

import (
	"github.com/xapier14/todo/internal/utils/jwt"
)

func VerifyAccessToken(token string) bool {
	valid, err := jwt.CheckIfValidJWT(token)
	if err != nil {
		return false
	}
	return valid
}