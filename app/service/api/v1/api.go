package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type API struct {
	MakeTransaction http.HandlerFunc
}

func (a *API) Routes(r *chi.Mux) {
	r.Post("/api/v1/transaction/{origin_document}/{destination_document}/{amount}", a.MakeTransaction)
}