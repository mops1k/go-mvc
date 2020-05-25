package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/handlers"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/config"
)

type Server struct {
	config      *config.ServerConfig
	srv         *http.Server
	middlewares []Middleware
	logger      *log.Logger
	routing     *routing
	handler     http.Handler
}

// Server constructor
func GetServer(config *config.ServerConfig, log *log.Logger) (s *Server) {
	s = &Server{
		config:  config,
		logger:  log,
		routing: Routing,
		handler: handlers.RecoveryHandler(handlers.RecoveryLogger(cli.Logger.Get(cli.ErrorLog).(handlers.RecoveryHandlerLogger)))(Routing.mux),
	}

	s.Middleware(&LoggingMiddleware{})

	return
}

// Start listening server
func (s *Server) ListenAndServe() error {
	s.routing.HandleControllers()
	http.Handle("/", s.routing.mux)
	for _, middleware := range s.middlewares {
		s.routing.mux.Use(middleware.Handler)
	}

	s.srv = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler:      s.handler,
		ReadTimeout:  s.getTimeout("read"),
		WriteTimeout: s.getTimeout("write"),
		IdleTimeout:  s.getTimeout("idle"),
		ErrorLog:     s.logger,
	}

	l, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		panic(err)
	}

	if s.config.CertFile != "" && s.config.KeyFile != "" {
		err := s.config.CheckTlsFiles()
		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}

		return s.srv.ServeTLS(l, s.config.CertFile, s.config.KeyFile)
	}

	return s.srv.Serve(l)
}

func (s *Server) Middleware(middleware Middleware) *Server {
	s.middlewares = append(s.middlewares, middleware)

	return s
}

func (s *Server) SetTimeouts(read uint16, write uint16, idle uint16) *Server {
	s.config.Timeouts["read"] = read
	s.config.Timeouts["write"] = write
	s.config.Timeouts["idle"] = idle

	return s
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) SetHandler(h http.Handler) *Server {
	s.handler = h

	return s
}

func (s *Server) String() string {
	protocol := "http"
	if s.config.CertFile != "" && s.config.KeyFile != "" {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d", protocol, s.config.Host, s.config.Port)
}

func (s *Server) getTimeout(name string) time.Duration {
	value, exists := s.config.Timeouts[name]
	if !exists {
		return 0
	}

	return time.Duration(value) * time.Second
}
