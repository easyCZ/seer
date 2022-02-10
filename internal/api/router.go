package api

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func AddRoutes(r *chi.Mux, svc *SyntheticsService) {
	r.Get("/synthetics", func(rw http.ResponseWriter, r *http.Request) {
		syns, err := svc.List(r.Context())
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(rw, r, http.StatusOK, syns); err != nil {
			log.Println("Failed to serialize synthetics to response")
		}
	})
}
