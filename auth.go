package main

import (
	// "fmt"
	"time"
	"encoding/json"
	"net/http"
)

type WebResponce struct {
	Success bool `json:"success"`
	Time int64 `json:"time"`
	User []User `json:"user"`
}

func (u *User) is_authenticated() bool {
	var sess Session
	var time = time.Now().Unix()
	// check session object is not expired
	return true
}

func GetHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var user User
	var resp WebResponce
	id := r.FormValue("id")
	q := a.db.Find(&user, id)
	var count int64
	q.Count(&count)
	if count == 0 {
		resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{}}
	} else {
		resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{user}}
	}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError, err
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func CreateHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, NotAllowedError
	}
	email := r.FormValue("email")
	username := r.FormValue("username")
	if email == "" || username == "" {
		return http.StatusBadRequest, InvalidRequestError
	}
	var (
		user User
		users []User
		success = false
	)
	a.db.Select("ID").Find(&users, "email=?", email)
	if len(users) == 0 {
		user = User{Email: email, Username: username, IsNew: true}
		db.Create(&user)
		success = true
	} else {
		user = users[0]
	}
	resp := ModelCreatedResponce{Success: success, ID: user.ID}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError, err
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func LoginHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusOK, nil
}
