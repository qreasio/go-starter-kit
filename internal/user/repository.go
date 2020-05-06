package user

import (
	"context"

	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/model"

	"github.com/jmoiron/sqlx"
)

var (
	listUsersSQL = "SELECT firstname,lastname,email,created,updated FROM users limit ?,?"
)

// Repository encapsulates the logic to user
type Repository interface {
	// Get returns the user with the specified user ID.
	// Get(ctx context.Context, id string) (*model.User, error)
	List(ctx context.Context, id *ListUsersRequest) ([]model.User, error)
}

// repository persists albums in database
type repository struct {
	db     *sqlx.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *sqlx.DB, log log.Logger) Repository {
	return repository{db: db, logger: log}
}

// List return slice of users base on request parameters
func (r repository) List(ctx context.Context, list *ListUsersRequest) ([]model.User, error) {
	users := make([]model.User, 0)
	offset := (list.Page - 1) * list.Limit
	err := r.db.Select(&users, listUsersSQL, offset, list.Limit)
	if err != nil {
		r.logger.Errorf("Failed to select users %s", err)
		return nil, err
	}
	return users, nil
}
