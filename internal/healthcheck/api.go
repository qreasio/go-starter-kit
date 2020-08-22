package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(r chi.Router) {
	r.Get("/health", Check)
}

// Check implements endpoint for health check
func Check(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "It is fine")
}
