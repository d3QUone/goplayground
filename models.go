package main

import (
	"time"
)

// NOTE: seems like gorm ignores 'not null' and other commands

type User struct {
	ID int `json:"id", orm:"auto"`
	Email string `json:"email", orm:"unique,size(255)"`
	Username string `json:"username", orm:"unique,size(255)"`
	IsNew bool `json:"is_new", orm:"bool"`
}

type Session struct {
	ID int `json:"id", orm:"auto"`
	User *User `json:"user", orm:"rel(fk)"`
	Code string `json:"code", orm:"unique,size(255)"`
	ExpiresAt time.Time `json:"expires_at", orm:"time.Time"`
}
