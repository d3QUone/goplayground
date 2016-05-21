package auth

import (
	"time"
	"encoding/json"
	"net/http"
)

type WebResponce struct {
	Success bool `json:"success"`
	Time time.Time `json:"time"`
	User User `json:"user"`
}

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
}

/*
func NewUser(name string) *DBConfig {
	var dbconfig = ReadConfig(name)
	return &dbconfig
}

func (s *DBConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json code error:", err)
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("json write error:", err)
	}
}
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := WebResponce{Success:true, Time:time.Now().UTC()}
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}
