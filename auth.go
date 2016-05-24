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

func (u *User) is_authenticated(a *AppContext) bool {
	// var sessions []Session
	var time = time.Now()
	rows, err := a.db.Query("SELECT s.*, u.* FROM `sessions` s JOIN `users` u ON s.user_id = u.id WHERE s.expires_at > $1 AND u.email = $2;", time, u.Email)
	fmt.Println("sessions =", rows)
	fmt.Println("err =", err)
	return true
}

func GetHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.FormValue("id")
	if id == "" {
		return http.StatusBadRequest, InvalidRequestError
	}
	resp := WebResponce{Success: true, Time: time.Now().Unix(), User: []User{}}
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
	var success = false
	var user_id = -1
	count, err := a.db.Query("SELECT COUNT(id) FROM users WHERE email = $1;", email)
	if GetRes(count) == 0 {
		res, err := a.db.Query("INSERT INTO users (email, username) VALUES ($1, $2) RETURNING id;", email, username)
		if err == nil {
			user_id = GetRes(res)
		}
	}
	resp := ModelCreatedResponce{Success: success, ID: user_id}
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
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, NotAllowedError
	}
	email := r.FormValue("email")
	if email == "" {
		return http.StatusBadRequest, InvalidRequestError
	}
	id, err := a.db.Query("SELECT id FROM users WHERE email = $1;", email)
	if err == nil {
		s_id, err := a.db.Query("INSERT INTO sessions (user_id, code, expires_at) VALUES ($1, $2, $3) RETURNING id;", GetRes(id), generate_code(SESSION_CODE_LEN), time.Now().Add(SESSION_DURATION))
		if err == nil {
			resp := ModelCreatedResponce{Success: true, ID: GetRes(s_id)}
			b, _ := json.Marshal(resp)
			w.Write(b)
			return http.StatusOK, nil
		}
	}
	return http.StatusInternalServerError, err
}

func GetLoggedInHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	// var time = time.Now()
	// var sessions []Session
	// var users []User
	// TODO: select related User Object
	// a.db.Where("expires_at>?", time).Find(&sessions).Association("UserID")
	// fmt.Printf("found %v sessions (%T):\n%v\n", len(sessions), sessions, sessions)
	// for i := 0; i < len(sessions); i++ {
	// 	s := sessions[i]
	// 	fmt.Printf("(%T)=%v\n", s, s)
	// 	// users = append(users, s.User)
	// }
	return http.StatusOK, nil
}
