package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// RegisterHealthRouter register routing for health check endpoint
func RegisterHealthRouter(r chi.Router) http.Handler {
	r.Get("/health", Check)
	return r
}

// Check implements endpoint for health check
func Check(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "It is fine")
}
