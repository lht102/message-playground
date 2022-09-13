package rest

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	jobworker "github.com/lht102/message-playground/jobworker"
	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 10 * time.Second
	shutdownTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *bunrouter.CompatRouter
	jobService jobworker.JobService
	logger     *zap.Logger
}

func NewServer(
	port int,
	jobService jobworker.JobService,
	logger *zap.Logger,
) *Server {
	srv := &Server{
		router:     bunrouter.New().Compat(),
		jobService: jobService,
		logger:     logger,
	}

	srv.routes()

	httpSrv := &http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           srv.router,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
	}

	srv.httpServer = httpSrv

	return srv
}

func (s *Server) Open() error {
	listener, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return fmt.Errorf("net listen on tcp: %w", err)
	}

	if err := s.httpServer.Serve(listener); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http listener serve: %w", err)
	}

	return nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http server shutdown: %w", err)
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
