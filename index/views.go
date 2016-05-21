package index

import (
	"encoding/json"
	"net/http"
)

type WebResponce struct {
	Success bool `json:"success"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := WebResponce{Success:true}
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}
