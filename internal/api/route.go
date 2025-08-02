package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/node/{id}", GetNodeHandler)
	r.Get("/subgraph/{id}", GetSubgraphHandler)

	return r
}
