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
	Age  int    `json:"age"`
}

func userHandler(ctx *gor.Context) {
	name, err := ctx.GetStringParam("name")
	if err != nil {
		ctx.WriteError(http.StatusBadRequest, err.Error())
		return
	}

	age, err := ctx.GetIntParam("age")
	if err != nil {
		ctx.WriteError(http.StatusBadRequest, err.Error())
		return
	}

	ctx.WriteJSON(http.StatusOK, response{Name: name, Age: age})
}

func main() {
	router := gor.NewRouter()
	router.GET(`/user/:name/:age`, userHandler)

	log.Fatalln(http.ListenAndServe(":9000", router))
}

```

```sh
$ curl -i "http://localhost:9000/user/John/23" 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 03 Mar 2021 09:56:53 GMT
Content-Length: 24

{"name":"John","age":23}
```
