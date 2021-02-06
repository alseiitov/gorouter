# gorouter

A simple HTTP router for Go.

Reimplemented [Custom HTTP Routing in Go](https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438) to write human-readable routes.


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
	"strconv"

	"github.com/alseiitov/gorouter"
)

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(ctx *gorouter.Context) {
	name := ctx.Params[0]
	age, err := strconv.Atoi(ctx.Params[1])

	if err != nil || age < 0 {
		ctx.WriteError(http.StatusBadRequest, "Invalid age")
		return
	}

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
Date: Sat, 06 Feb 2021 13:23:52 GMT
Content-Length: 24

{"name":"John","age":23}
```
```sh
$ curl -X POST -i "http://localhost:9000/user/John/23" 
HTTP/1.1 405 Method Not Allowed
Content-Type: application/json
Date: Sat, 06 Feb 2021 13:25:33 GMT
Content-Length: 30

{"error":"Method not allowed"}
```