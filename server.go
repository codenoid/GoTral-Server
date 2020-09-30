package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mervick/aes-everywhere/go/aes256"
)

var (
	configPath      string
	encryptPassword = "somehardpw"

	// BasicAuth functionality
	BasicAuth  bool
	BAUsername string
	BAPassword string

	// server utility
	bind string
)

// getConfig : handler for getting config data
func getConfig(w http.ResponseWriter, r *http.Request) {

	if BasicAuth == true {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != BAUsername && password != BAPassword {
			http.Error(w, "Not authorized", 401)
			return
		}
	}

	// filename of project id
	configFile := r.URL.Query().Get("id")

	// path already end with slash
	path := configPath + "/" + configFile
	if fileExists(path) {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			http.Error(w, "Failed to read file", 500)
			return
		}

		encrypted := aes256.Encrypt(string(b), encryptPassword)

		// return json content from file as string
		w.Write([]byte(encrypted))
		return
	}

	http.Error(w, "File doesn't exist!", 404)
}

func main() {

	flag.StringVar(&bind, "bind", ":8080", "-bind 0.0.0.0:3000")
	flag.Parse()

	http.HandleFunc("/config", getConfig)

	fmt.Println("starting to listen on", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
