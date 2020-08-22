package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"

	e "github.com/qreasio/go-starter-kit/pkg/error"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/mid"
	"github.com/qreasio/go-starter-kit/pkg/model"

	"github.com/go-chi/render"
)

// HTTP defines user http handlers interface
type HTTP interface {
	List(w http.ResponseWriter, r *http.Request)
}

// UserHTTP is struct to represent user http transport
type userHTTP struct {
	svc    Service
	Logger log.Logger
}

// NewUserHTTP returns ne UserHTTP struct instance
func NewUserHTTP(db *sqlx.DB, logger log.Logger, validator *validator.Validate) HTTP {
	repo := NewRepository(db, logger)
	svc := NewService(repo, validator, logger)
	return userHTTP{svc: svc, Logger: logger}
}

// GetUserHTTP returns ne UserHTTP struct instance
func GetUserHTTP(svc Service, logger log.Logger) HTTP {
	return userHTTP{svc: svc, Logger: logger}
}

// listRequestDecoder decode http request to construct ListUsersRequest
func listRequestDecoder(r *http.Request) *ListUsersRequest {
	pagination, ok := mid.PaginateFromContext(r.Context())
	if !ok {
		pagination = model.NewPagination()
	}
	search := r.URL.Query().Get("search")
	listRequest := &ListUsersRequest{Pagination: *pagination, Search: search}
	return listRequest
}

// List is handler for list users endpoint
func (h userHTTP) List(w http.ResponseWriter, r *http.Request) {
	listRequest := listRequestDecoder(r)
	users, err := h.svc.ListUsers(r.Context(), listRequest)
	if err != nil {
		h.Logger.With(r.Context()).Errorf("list users error : %s", err)
		render.Render(w, r, e.BadRequest(err, "bad request"))
		return
	}
	render.Respond(w, r, users)
}
