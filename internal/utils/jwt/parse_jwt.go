package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(getKey()), nil
}

func ParseJWT(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, keyFunc)
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}