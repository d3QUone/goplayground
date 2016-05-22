package main

import (
	"time"
)

// NOTE: seems like gorm ignores 'not null' and other commands

type User struct {
	ID uint `json:"id", gorm:"not null;unique"`
	Email string `json:"email", gorm:"not null;unique"`
	Username string `json:"username", gorm:"not null;unique"`
	IsNew bool `json:"is_new"`
}

type Session struct {
	ID uint `json:"id", gorm:"not null;unique"`
	User User `gorm:"ForeignKey:UserID"`
	UserID uint
	Code string `json:"code", gorm:"unique"`
	ExpiresAt time.Time `json:"expires_at"`
}
