package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"goplayground/index"
	"goplayground/auth"
)

// ================== //
// Config preparation //
// ================== //

var PORT = flag.String("port", ":8080", "Listen address")

const DB_CONFIG = "config.json"

type DBConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func ReadConfig(configfile string) (configuration DBConfig) {
	_, err := os.Stat(configfile)
	if err != nil {
		fmt.Printf("Config file '%v' is missing!\n", configfile)
		panic(err)
	}
	file, err :=  ioutil.ReadFile(configfile)  // os.Open(configfile)
	if err != nil {
		fmt.Printf("file not found (%v)\n", err)
		panic(err)
	}
	json.Unmarshal(file, &configuration)
	return
}

// ==== //
// Main //
// ==== //

func main() {
	// config := ReadConfig(DB_CONFIG)

	// routes
	http.HandleFunc("/", index.Handler)
	http.HandleFunc("/auth/", auth.Handler)

	log.Fatal(http.ListenAndServe(*PORT, nil))
}
