package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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


type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}


func main() {
	var dbconfig = ReadConfig("config.json") 
	fmt.Printf("%T\n%v\n", dbconfig, dbconfig)
	fmt.Printf("name=%v\n", dbconfig.Name)

	// b, err := json.Marshal(dbconfig)
	// if err != nil {
	// 	fmt.Println("json code error:", err)
	// }
	// fmt.Println(b)

}
