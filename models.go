package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	Username string `json:"username"`
}
