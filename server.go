package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

import "github.com/mervick/aes-everywhere/go/aes256"

var (
	path            = "./config/" // remember to end with /
	encryptPassword = "somehardpw"
	BasicAuth       = false
	validUsername   = "guest"
	validPassword   = "guest"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// getConfig : handler for getting config data
func getConfig(w http.ResponseWriter, r *http.Request) {

	if BasicAuth == true {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != validUsername || password != validPassword {
			http.Error(w, "Not authorized", 401)
			return
		}
	}

	if r.Method == "GET" {
		// get url query
		keys := r.URL.Query()

		// filename of project id
		ConfigFile := keys.Get("id")

		// path already end with slash
		if fileExists(path + ConfigFile) {
			b, err := ioutil.ReadFile(path + ConfigFile)
			if err != nil {
				http.Error(w, "Failed to read file", 500)
				return
			}

			encrypted := aes256.Encrypt(string(b), encryptPassword)

			// return json content from file as string
			w.Write([]byte(encrypted))
			return
		} else {
			http.Error(w, "File doesn't exist!", 500)
			return
		}
	}
	fmt.Fprintf(w, "only get my dudeeeeeeeeeeee")
}

func main() {
	http.HandleFunc("/config", getConfig)
	fmt.Println("starting to listen on :6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
