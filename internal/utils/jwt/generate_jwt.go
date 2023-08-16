package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// The default number of minutes until the JWT token expires.
	DefaultExpiration = 15
)

var key *string = nil

func getKey() string {
	if key == nil {
		envKey := os.Getenv("JWT_KEY")
		if envKey == "" {
			panic("JWT_KEY not set")
		}
		key = &envKey
	}
	return *key
}

// Generates a new JWT token. This token does not expire.
func GenerateJWT(claims map[string]interface{}) (string, error) {
	key := getKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString([]byte(key))
}

// Generates a new JWT token. This token expires after the given number of minutes.
func GenerateExpiringJWT(claims map[string]interface{}, minutes int) (string, error) {
	SetExpirationTime(claims, minutes)
	token, err := GenerateJWT(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Generates a new JWT token. This token expires after the default number of minutes given by jwt.DefaultExpiration.
func GenerateDefaultExpiringJWT(claims map[string]interface{}) (string, error) {
	return GenerateExpiringJWT(claims, DefaultExpiration)
}

// Checks if the given JWT token is valid and not expired.
func CheckIfValidJWT(token string) (bool, error) {
	key := getKey()
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false, err
	}
	return t.Valid, nil
}