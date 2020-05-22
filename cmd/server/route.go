package server

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/qreasio/go-starter-kit/internal/healthcheck"
	usertransport "github.com/qreasio/go-starter-kit/internal/user/transport"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/mid"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Routing setup api routing
func Routing(db *sqlx.DB, logger log.Logger) chi.Router {
	validate = validator.New()
	// setup server routing

	r := chi.NewRouter()

	// homepage welcome page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.HTML(w, r, "<html><head><title>Go Starter Kit</title></head><body>Welcome to Go Starter Kit</head></body></html>")
	})

	// register health check route
	healthcheck.RegisterHealthRouter(r)

	r.Route("/v1", func(r chi.Router) {
		r.Use(mid.APIVersionCtx("v1"))
		r.Mount("/users", usertransport.RegisterUserRouter(usertransport.NewUserHTTP(db, logger, validate)))
	})

	return r
}
