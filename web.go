package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"goplayground/auth"
)

// ================== //
// Config preparation //
// ================== //

var PORT = flag.String("port", ":8080", "Listen address")

const DB_CONFIG = "config.json"

type DBConfig struct {
	Host 	string 	`json:"host"`
	Port 	int		`json:"port"`
	Name 	string 	`json:"name"`
	User 	string 	`json:"user"`
	Pass 	string 	`json:"pass"`
}

func ReadConfig(configfile string) (configuration DBConfig) {
	_, err := os.Stat(configfile)
	if err != nil {
		fmt.Println("Config file is missing: ", configfile)
	}
	file, err :=  ioutil.ReadFile(configfile)  // os.Open(configfile)
	if err != nil {
		fmt.Println("file not found (", err, ")")
	}
    json.Unmarshal(file, &configuration)
	return
}


// ======== //
// Core app //
// ======== //

func NewConfig(name string) *DBConfig {
	var dbconfig = ReadConfig(DB_CONFIG)
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

func main() {
	fmt.Printf("F=(%T)%v\n", auth.F(), auth.F())	

	// routes
	http.Handle("/", NewConfig(DB_CONFIG))

	log.Fatal(http.ListenAndServe(*PORT, nil))
}
