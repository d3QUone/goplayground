package main

import (
	"encoding/json"
	"net/http"
)

type IndexWebResponce struct {
	Success bool `json:"success"`
}

func Handler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	resp := IndexWebResponce{Success: true}
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
	return http.StatusOK, nil
}

func SecretHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {() {
	return http.StatusOK, nil
}
