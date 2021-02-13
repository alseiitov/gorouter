package gorouter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	http.ResponseWriter
	*http.Request
	Params map[string]string
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

func (ctx *Context) ReadBody(data interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) setURLValues(keys, values []string) {
	for i, key := range keys {
		ctx.Set(key, values[i])
	}
}

func (ctx *Context) Set(key, value string) {
	ctx.Params[key] = value
}
