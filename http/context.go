package http

import "net/http"

type Context struct {
	response   http.ResponseWriter
	request    *http.Request
	statusCode int
	headers    map[string]string
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
