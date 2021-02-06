package gorouter

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

type Error struct {
	Error string `json:"error"`
}

func (ctx *Context) WriteString(code int, body string) {
	ctx.ResponseWriter.Header().Set("Content-Type", "text/plain")
	ctx.WriteHeader(code)

	ctx.ResponseWriter.Write([]byte(body))
}

func (ctx *Context) WriteJSON(code int, body []byte) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.WriteHeader(code)

	ctx.ResponseWriter.Write(body)
}

func (ctx *Context) WriteError(code int, err string) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.WriteHeader(code)

	jsonData, _ := json.Marshal(&Error{Error: err})
	ctx.ResponseWriter.Write(jsonData)
}
