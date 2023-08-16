package jwt

import (
	"time"
)

// Sets the expiration time of the JWT token.
func SetExpirationTime(claimsMap map[string]interface{}, minutes int) time.Time {
	expirationTime := time.Now().Add(time.Minute * time.Duration(minutes))
	claimsMap["exp"] = expirationTime.Unix()
	return expirationTime
}