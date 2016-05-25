package main

import (
	"fmt"
	"time"
)

// TODO: create BaseModel interface with TableName and 'schema()'

// ==== USER ====

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	IsNew    bool   `json:"is_new"`
}

func (o *User) NeedMigration() bool {
	return false
}

func (o *User) TableName() string {
	return "users"
}

func (o *User) Schema() string {
	tn := o.TableName()
	return fmt.Sprintf(`
DROP TABLE IF EXISTS %v;
CREATE TABLE %v (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL UNIQUE,
	username VARCHAR(255) NOT NULL UNIQUE,
	is_new BOOLEAN NOT NULL DEFAULT TRUE
);`, tn, tn)
}

// ==== SESSION ====

type Session struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (o *Session) NeedMigration() bool {
	return false
}

func (o *Session) TableName() string {
	return "sessions"
}

func (o *Session) Schema() string {
	tn := o.TableName()
	return fmt.Sprintf(`
DROP TABLE IF EXISTS %v;
CREATE TABLE %v (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users ON DELETE RESTRICT,
	code VARCHAR(255) NOT NULL UNIQUE,
	expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);`, tn, tn)
}
