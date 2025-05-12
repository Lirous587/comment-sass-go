package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"size:32;not null"`
	Email        string `gorm:"uniqueIndex;not null;size:255"`
	HashPassword []byte `gorm:"type:varchar(255)"`
	GithubID     int    `gorm:"uniqueIndex;type:int;not null"`
}

type JwtPayload struct {
	ID        uint      `json:"id"`
	LoginType LoginType `json:"login_type"`
}

type RefreshTokenReq = JwtPayload
type RefreshTokenRes struct {
	Token string `json:"token"`
}
