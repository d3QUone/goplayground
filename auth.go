package main

/*
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
	var sessions []Session
	var time = time.Now()
	// check session object is not expired
	a.db.Where("expires_at>?", time).Order("expires_at").Find(&sessions)
	fmt.Println("sessions =", sessions)
	return true
}

func GetHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var user User
	var resp WebResponce
	id := r.FormValue("id")
	a.db.Find(&user, id)
	fmt.Printf("users (%T)\n", user)
	
	// TODO: fix "get" user

	// if count == 0 {
	// 	resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{}}
	// } else {
		resp = WebResponce{Success: true, Time: time.Now().Unix(), User: []User{user}}
	// }
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
	var user User
	var users []User
	var success = false
	a.db.Select("ID").Where("email=?", email).Find(&users)
	if len(users) == 0 {
		user = User{Email: email, Username: username, IsNew: true}
		a.db.Create(&user)
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
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, NotAllowedError
	}
	email := r.FormValue("email")
	if email == "" {
		return http.StatusBadRequest, InvalidRequestError
	}
	// load user
	var user User
	var users []User
	a.db.Where("email=?", email).Find(&users)
	if len(users) == 1 {
		user = users[0]
	} else {
		return http.StatusInternalServerError, NotAuthorizedError
	}
	// create new session
	session := Session{User: &user, Code: generate_code(SESSION_CODE_LEN), ExpiresAt: time.Now().Add(SESSION_DURATION)}
	a.db.Create(&session)
	return http.StatusOK, nil
}

func GetLoggedInHandler(a *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var time = time.Now()
	var sessions []Session
	// var users []User
	// TODO: select related User Object
	a.db.Where("expires_at>?", time).Find(&sessions).Association("UserID")
	fmt.Printf("found %v sessions (%T):\n%v\n", len(sessions), sessions, sessions)
	for i := 0; i < len(sessions); i++ {
		s := sessions[i]
		fmt.Printf("(%T)=%v\n", s, s)
		// users = append(users, s.User)
	}
	return http.StatusOK, nil
}
*/
