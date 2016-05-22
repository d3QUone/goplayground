package main

import (
	"errors"
)

var (
	NotAllowedError = errors.New("Method not allowed")
	InvalidRequestError = errors.New("Invalid request")
)

type ModelCreatedResponce struct {
	Success bool `json:"success"`
	ID uint `json:"id"`
}

func generate_code() {
	
}
