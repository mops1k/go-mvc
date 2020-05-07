package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	port        int
	host        string
	timeouts    map[string]uint16
	srv         *http.Server
	middlewares []Middleware
	logger      *log.Logger
	routing     *Routing
}

func GetServer(host string, port int, log *log.Logger) (s *Server) {
	s = &Server{
		port:     port,
		host:     host,
		timeouts: make(map[string]uint16),
		logger:   log,
		routing:  GetRouting(),
	}
	s.Middleware(&LoggingMiddleware{})

	return
}

func (s *Server) ListenAndServe() error {
	s.routing.HandleControllers()
	http.Handle("/", s.routing.mux)
	for _, middleware := range s.middlewares {
		s.routing.mux.Use(middleware.Handler)
	}

	s.srv = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:      s.routing.mux,
		ReadTimeout:  s.getTimeout("read"),
		WriteTimeout: s.getTimeout("write"),
		IdleTimeout:  s.getTimeout("idle"),
		ErrorLog:     s.logger,
	}

	l, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		panic(err)
	}

	return s.srv.Serve(l)
}

func (s *Server) Middleware(middleware Middleware) *Server {
	s.middlewares = append(s.middlewares, middleware)

	return s
}

func (s *Server) SetTimeouts(read uint16, write uint16, idle uint16) *Server {
	s.timeouts["read"] = read
	s.timeouts["write"] = write
	s.timeouts["idle"] = idle

	return s
}

func (s *Server) String() string {
	return fmt.Sprintf("http://%s:%d", s.host, s.port)
}

func (s *Server) getTimeout(name string) time.Duration {
	value, exists := s.timeouts[name]
	if !exists {
		return 0
	}

	return time.Duration(value) * time.Second
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
