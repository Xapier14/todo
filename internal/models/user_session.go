package models

import (
	"time"
)

type SessionInfo struct {
	IP string `json:"ip"`
	UserAgent string `json:"userAgent"`
}

type UserSession struct {
	ID        	uint		`json:"id" gorm:"primaryKey"`
	UserID    	uint		`json:"userId"`
	SessionInfo SessionInfo	`json:"sessionInfo" gorm:"embedded"`
	Token     	string		`json:"token" gorm:"unique"`
	ExpiresAt 	time.Time	`json:"expiresAt"`
}