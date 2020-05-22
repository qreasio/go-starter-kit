package transport

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qreasio/go-starter-kit/internal/user"
	e "github.com/qreasio/go-starter-kit/pkg/error"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/mid"
	"github.com/qreasio/go-starter-kit/pkg/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

// UserHTTP is struct to represent user http transport
type UserHTTP struct {
	svc    user.Service
	logger log.Logger
}

// NewUserHTTP returns ne UserHTTP struct instance
func NewUserHTTP(db *sqlx.DB, logger log.Logger, validator *validator.Validate) UserHTTP {
	repo := user.NewRepository(db, logger)
	svc := user.NewService(repo, validator, logger)
	return UserHTTP{svc: svc, logger: logger}
}

// RegisterUserRouter registers router for users endpoint
func RegisterUserRouter(http UserHTTP) http.Handler {
	r := chi.NewRouter()
	r.With(mid.Paginate).Get("/", http.List)
	return r
}

func listRequestDecoder(r *http.Request) *user.ListUsersRequest {
	pagination, ok := mid.PaginateFromContext(r.Context())
	if !ok {
		pagination = model.NewPagination()
	}
	search := r.URL.Query().Get("search")
	listRequest := &user.ListUsersRequest{Pagination: *pagination, Search: search}
	return listRequest
}

// List is handler for list users endpoint
func (h UserHTTP) List(w http.ResponseWriter, r *http.Request) {
	listRequest := listRequestDecoder(r)
	users, err := h.svc.ListUsers(r.Context(), listRequest)
	if err != nil {
		h.logger.With(r.Context()).Errorf("list users error : %s", err)
		render.Render(w, r, e.BadRequest(err, "bad request"))
		return
	}
	render.Respond(w, r, users)
}
