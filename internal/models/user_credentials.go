package models

type UserCredential struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"passwordHash"`
	Salt         string `json:"salt"`
}