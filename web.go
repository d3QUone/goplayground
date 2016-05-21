package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"flag"
	"net/http"
)

//// Config preparation ////

var port = flag.String("port", ":8080", "Listen address")

const DB_CONFIG = "config.json"

type DBConfig struct {
	Name string
	Host string
	Port string
	User string
	Pass string
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


//// Web app ////

func NewConfig(name string) *DBConfig {
	var dbconfig = ReadConfig(DB_CONFIG) 
	return &dbconfig
}

func (s *DBConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json code error:", err)
	}
	c, err := w.Write(b)
	if err != nil {
		fmt.Println("json write error:", err)
	}
	fmt.Println("c =", c)
}

func main() {
	// var dbconfig = ReadConfig(DB_CONFIG) 
	// fmt.Printf("%T\n%v\n", dbconfig, dbconfig)

	// routes
	http.Handle("/", NewConfig(DB_CONFIG))

	http.ListenAndServe(*port, nil)	
}
