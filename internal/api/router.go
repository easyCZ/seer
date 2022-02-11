package api

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func AddRoutes(r *chi.Mux, svc *SyntheticsService) {
	r.Get("/synthetics", func(rw http.ResponseWriter, r *http.Request) {
		syns, err := svc.List(r.Context())
		if err != nil {
			log.Println("Failed to list synthetics", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(rw, r, http.StatusOK, syns); err != nil {
			log.Println("Failed to write list of synthetics to response", err)
		}
	})

	r.Get("/synthetics/{id}", func(w http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(rawID, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		s, err := svc.Get(r.Context(), uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(w, r, http.StatusOK, s); err != nil {
			log.Println("Failed to write synthetic to response", err)
		}
	})
}
