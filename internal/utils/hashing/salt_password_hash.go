package hashing

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return string(hash)
}

func GenerateSalt() string {
	varSize, err := rand.Int(rand.Reader, big.NewInt(8))
	if err != nil {
		panic(err)
	}
	size := varSize.Int64() + 8
	salt := make([]byte, size)
	_, err = rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return string(salt)
}