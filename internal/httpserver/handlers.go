package httpserver

import (
	"fmt"
	"net/http"
	"encoding/json"

	"chi_nats/internal/messaging"

	"github.com/go-chi/chi"
)

func publishMessage(
	p broker.Publisher,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		subject := chi.URLParam(r, "subjectName")

		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Failed to parse request body: %v", err),
				http.StatusBadRequest,
			)
		}

		p.PublishMsg(subject, []byte(requestBody.Message))
	}
}
