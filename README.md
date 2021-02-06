# gorouter

A simple HTTP router for Go.

Reimplemented [Custom HTTP Routing in Go](https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438) to write human-readable routes and to easily parse values from URI.


### Install
```sh
$ go get github.com/alseiitov/gorouter
```


### Example
```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alseiitov/gorouter"
)

type data struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func userHandler(ctx *gorouter.Context) {
	name := ctx.Params["name"]
	age := ctx.Params["age"]

	jsonData, _ := json.Marshal(&data{Name: name, Age: age})
	ctx.WriteJSON(http.StatusOK, jsonData)
}

func main() {
	router := gorouter.NewRouter()
	router.GET(`/user/:name/:age`, userHandler)

	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatalln(err.Error())
	}
}

```

```sh
$ curl -i "http://localhost:9000/user/John/23" 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 06 Feb 2021 18:10:29 GMT
Content-Length: 26

{"name":"John","age":"23"}
```
