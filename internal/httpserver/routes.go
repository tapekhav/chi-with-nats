package httpserver

import (
	"net/http"

	"chi_nats/internal/messaging"

	"github.com/go-chi/chi"
)

func SetRoutes(r *chi.Mux, p broker.Publisher) {
	r.MethodFunc(
		"GET", 
		"/ping", 
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		},
	)

	r.Post("/publish/{subjectName}", publishMessage(p))
}
