package rest

import (
	"context"
	"errors"
	"net/http"

	"get-time/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log     *logrus.Entry
	httpSrv *http.Server
	service service.Checker
	host    string
}

// Run server.
func (s *Server) Run(ctx context.Context) {
	go func() {
		s.log.Infof("Listening on %s", s.host)
		if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Panic(err)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := s.httpSrv.Shutdown(ctx); err != nil {
			s.log.Info("Closing HTTP server", err)
		}
	}()
	return
}

func NewServer(log *logrus.Logger, service service.Checker, addr string) *Server {
	srv := Server{
		log:     log.WithField("module", "server"),
		service: service,
		host:    "localhost" + addr,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/api/v1/time", srv.handlerGetTime)

	server := http.Server{
		Addr:    srv.host,
		Handler: r,
	}
	srv.httpSrv = &server

	return &srv
}
