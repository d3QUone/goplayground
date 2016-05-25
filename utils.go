package main

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"
)

// ==== DATABASE HELPERS ====

func GetRes(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	return count
}

// ==== ERRORS ====

var (
	NotAllowedError     = errors.New("Method not allowed")
	InvalidRequestError = errors.New("Invalid request")
	NotAuthorizedError  = errors.New("Not authorized")
)

type ModelCreatedResponce struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
}

type ErrorResponce struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// huge thanks to http://stackoverflow.com/a/31832326/4213969

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generate_code(n int) string {
	b := make([]byte, n)
	src := rand.NewSource(time.Now().UnixNano())
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
