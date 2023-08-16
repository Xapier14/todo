package hashing

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID.
func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}