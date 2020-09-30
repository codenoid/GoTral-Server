# GoTral Server

Server for [gotral library](https://github.com/codenoid/GoTral)

## Adding Your Config Data

1. Read server.go and config folder

## Usage

1. clone and run the server !
```
$ git clone  https://github.com/codenoid/GoTral-Server.git
$ cd GoTral-Server
// edit some config / do what you want
$ go run server.go
starting to listen on :8080
```
2. accessing the data

- http://localhost:8080/config?id=filename.ext
- just like http://localhost:8080/config?id=ecommerce.json
- ecommerce.json must be inside [choosed folder](https://github.com/codenoid/GoTral-Server/blob/ca0c016c2642ab91d27ea8369a74cb9818d94f79/server.go#L14) that contain config file

```
package main

import (
	"fmt"

	"github.com/codenoid/gotral"
)

func main() {

	// super secret key
	secret := "somehardpw" // or just put string in there

	config, err := gotral.DirectLoad("http://localhost:8080/config?id=ecommerce.json", secret)
	if err != nil { fmt.Println(err) }
	if val, err := config.Get("mysql_username"); !err {
 		fmt.Println(val)
 	}

	// with basic auth support
	withOpt := gotral.GoTral{
		Url: "http://localhost:8080/config?id=ecommerce.json",
		Passphrase: "somehardpw",
		BasicAuth: true,
		Username: "guest",
		Password: "guest",
	}

	config, err = withOpt.LoadConfig()
	if err != nil { fmt.Println(err) }
	if val, err := config.Get("mysql_username"); !err {
 		fmt.Println(val)
 	}
}
```
