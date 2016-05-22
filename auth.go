package main

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
)

type WebResponce struct {
	Success bool `json:"success"`
	Time int64 `json:"time"`
	User []User `json:"user"`
}

// TODO: move to utils, into UnifyCreateResponce (if needed)
type UserCreateResponce struct {
	Success bool `json:"success"`
	ID uint `json:"id"`
}

func GetHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var user User
	id := r.FormValue("id")
	q := a.db.Find(&user, id)
	var count int64
	q.Count(&count)
	// fmt.Println("user =", user)
	// fmt.Println("count =", count)
	var resp WebResponce
	if count == 0 {
		resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{}}
	} else {
		resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{user}}
	}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 502, err
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 503, err
	}
	return 200, nil
}


func CreateHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, NotAllowedError
	}
	email := r.FormValue("email")
	username := r.FormValue("username")

	var (
		user User 
		count int64
		success = false
	)
	// var 
	q := a.db.First(&user, "email = ?", email)
	q.Count(&count)
	if count == 0 {
		user = User{Email: email, Username: username}
		db.Create(&user)
		success = true
	}
	fmt.Println("user=", user, "count=", count)
	resp := UserCreateResponce{Success: success, ID: user.ID}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 502, err
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 503, err
	}
	return 200, nil	
}
