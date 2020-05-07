package http

import (
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/mops1k/go-mvc/cli"
)

type Middleware interface {
	Handler(next http.Handler) http.Handler
}

type LoggingMiddleware struct {
}

func (lm *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(cli.Logger.CreateLogFile("access.log"), next)
}
