package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}
