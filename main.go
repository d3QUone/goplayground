package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	PORT   = flag.String("port", ":8080", "Listen address")
	config DBConfig
	db     *sql.DB
)

const (
	DB_CONFIG         = "config.json"
	SESSION_DURATION  = time.Minute * 10
	SESSION_MIN_RENEW = time.Minute * 2 // create new session if the latest expires in 'SESSION_MIN_RENEW'
	SESSION_CODE_LEN  = 45
)

type AppContext struct {
	db *sql.DB
}

type AppHandler struct {
	*AppContext
	H func(*AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Updated to pass ah.AppContext as a parameter to our handler type.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	status, err := ah.H(ah.AppContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
			// And if we wanted a friendlier error page, we can now leverage our context instance - e.g.
			// err := ah.renderTemplate(w, "http_404.tmpl", nil)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

type DBConfig struct {
	Provider string `json:"provider"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}

func (c *DBConfig) get() string {
	if len(c.Pass) > 0 {
		return "host=" + c.Host + " port=" + c.Port + " dbname=" + c.Name + " sslmode=disable user=" + c.User + " password=" + c.Pass
	} else {
		return "host=" + c.Host + " port=" + c.Port + " dbname=" + c.Name + " sslmode=disable user=" + c.User
	}
}

func ReadConfig(configfile string) (configuration DBConfig) {
	_, err := os.Stat(configfile)
	if err != nil {
		fmt.Printf("Config file '%v' is missing!\n", configfile)
		panic(err)
	}
	file, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Printf("file not found (%v)\n", err)
		panic(err)
	}
	json.Unmarshal(file, &configuration)
	return
}

func init() {
	config = ReadConfig(DB_CONFIG)
	fmt.Println(config.get())
	var err error
	db, err = sql.Open(config.Provider, config.get())
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// migrations
	var u User
	if u.NeedMigration() {
		_, err := db.Exec(u.Schema())
		if err != nil {
			panic(err)
		}
	}
	var s Session
	if s.NeedMigration() {
		_, err := db.Exec(s.Schema())
		if err != nil {
			panic(err)
		}
	}
}

// ======
// Routes
// ======

func main() {
	context := &AppContext{db: db}

	http.Handle("/", AppHandler{context, Handler})
	http.Handle("/auth/get", AppHandler{context, GetHandler})
	http.Handle("/auth/create", AppHandler{context, CreateHandler})
	http.Handle("/auth/login", AppHandler{context, LoginHandler})
	http.Handle("/auth/logged", AppHandler{context, GetLoggedInHandler})

	log.Fatal(http.ListenAndServe(*PORT, nil))
}
