package httpserver

import (
	"net/http"

	"chi_nats/internal/messaging"

	"github.com/go-chi/chi"
)

func handleAllMethods(
	pattern string,
	handler http.HandlerFunc,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.MethodFunc(http.MethodGet, pattern, handler)
		r.MethodFunc(http.MethodPost, pattern, handler)
		r.MethodFunc(http.MethodDelete, pattern, handler)
		r.MethodFunc(http.MethodPut, pattern, handler)
		r.MethodFunc(http.MethodHead, pattern, handler)
	}
}

func SetRoutes(r *chi.Mux, p broker.Publisher) {
	r.Route(
		"/ping",
		handleAllMethods(
			"/", 
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			},
		),
	)

	r.Route("/publish", func(apiRoute chi.Router) {
		apiRoute.Post("/{subjectName}", publishMessage(p))
	})
}
