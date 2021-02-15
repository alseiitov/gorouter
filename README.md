# gorouter

A simple HTTP router for Go.

Reimplemented [Custom HTTP Routing in Go](https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438) to write human-readable routes and to easily parse values from URI.


### Install
```sh
$ go get -u github.com/alseiitov/gorouter
```


### Example
```go
package main

import (
	"log"
	"net/http"

	"github.com/alseiitov/gorouter"
)

type response struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func userHandler(ctx *gorouter.Context) {
	name, _ := ctx.GetParam("name")
	age, _ := ctx.GetParam("age")

	err := ctx.WriteJSON(http.StatusOK, response{Name: name, Age: age})
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	router := gorouter.NewRouter()
	router.GET(`/user/:name/:age`, userHandler)

	log.Fatalln(http.ListenAndServe(":9000", router))
}

```

```sh
$ curl -i "http://localhost:9000/user/John/23" 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 14 Feb 2021 17:03:22 GMT
Content-Length: 27

{"name":"John","age":"23"}
```
