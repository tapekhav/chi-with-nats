package httpserver

import (
	"context"
	"net/http"

	"chi_nats/internal/messaging"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	httpServer *http.Server
}

func New(url string, p broker.Publisher) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	s := Server{
		httpServer: &http.Server{
			Addr: url,
			Handler: r,
		},
	}

	SetRoutes(r, p)

	return &s
}

func (s *Server) Listen(ctx context.Context) error {
	errorChan := make(chan error)

	defer func() {
		close(errorChan)
	}()

	go func() {
		errorChan <- s.httpServer.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errorChan:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}
