package http

import (
	"log"
	"net/http"

	"github.com/mops1k/go-mvc/cli"
)

type Context struct {
	response   http.ResponseWriter
	request    *http.Request
	statusCode int
	headers    map[string]string
	vars       map[string]string
}

func (c *Context) Get(key string) string {
	if _, exists := c.vars[key]; !exists {
		cli.Logger.Get(cli.AppLog).(*log.Logger).Printf(`Path var "%s" does no exists`, key)
	}

	return c.vars[key]
}

func (c *Context) Response() http.ResponseWriter {
	return c.response
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) StatusCode(code int) {
	c.statusCode = code
}

func (c *Context) Header(name string, value string) {
	c.headers[name] = value
}

func (c *Context) Headers() map[string]string {
	return c.headers
}
