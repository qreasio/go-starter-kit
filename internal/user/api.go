package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/mid"
)

// RegisterHandlers registers handlers for specified path
func RegisterHandlers(r chi.Router, db *sqlx.DB, logger log.Logger, validate *validator.Validate) {
	r.Mount("/users", RegisterHTTPHandlers(NewUserHTTP(db, logger, validate)))
}

// RegisterHTTPHandlers registers http handlers for users endpoint
func RegisterHTTPHandlers(http HTTP) http.Handler {
	r := chi.NewRouter()
	r.With(mid.Paginate).Get("/", http.List)
	return r
}
