package rest

import (
	"context"
	"get-time/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Server struct {
	httpSrv *http.Server
	service service.CheckerI
}

// Run server.
func (s *Server) Run(ctx context.Context) error {
	go func() {
		log.Println("Run")
		if err := s.httpSrv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := s.httpSrv.Shutdown(ctx); err != nil {
			log.Println("Closing HTTP server", err)
		}
	}()
	return nil
}

func NewServer(service service.CheckerI) *Server {
	srv := Server{
		httpSrv: nil,
		service: service,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/gettime", srv.HandlerGetTime)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	srv.httpSrv = &server

	return &srv
}
