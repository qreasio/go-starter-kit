package mid

import (
	"context"
	"net/http"
	"strconv"

	e "github.com/qreasio/go-starter-kit/pkg/error"
	"github.com/qreasio/go-starter-kit/pkg/model"

	"github.com/go-chi/render"
)

// Paginate is middleware to construct Pagination from http.Request
func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagination := model.NewPagination()
		sortParam := r.URL.Query().Get("sort")
		if sortParam != "" {
			pagination.Sort = sortParam
		}
		page := r.URL.Query().Get("page")
		if page != "" {
			p, err := strconv.Atoi(page)
			if err != nil {
				render.Render(w, r, e.InvalidParameter(err, "invalid Page parameter"))
				return
			}
			pagination.Page = p
		}

		limitParam := r.URL.Query().Get("limit")
		if limitParam != "" {
			limit, err := strconv.Atoi(limitParam)

			if err != nil {
				render.Render(w, r, e.InvalidParameter(err, "invalid Limit parameter"))
				return
			}
			pagination.Limit = limit
		}

		r = r.WithContext(context.WithValue(r.Context(), paginationKey, pagination))
		next.ServeHTTP(w, r)
	})
}

// PaginateFromContext get Paginate from context
func PaginateFromContext(ctx context.Context) (*model.Pagination, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	userIP, ok := ctx.Value(paginationKey).(*model.Pagination)
	return userIP, ok
}
