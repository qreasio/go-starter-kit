package user

import (
	"context"

	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/model"

	"github.com/jmoiron/sqlx"
)

var (
	// ListUsersSQL is SQL Clause to select public users data
	ListUsersSQL = "SELECT id, username, first_name, last_name, email, date_joined, last_login, is_active, is_staff, is_superuser FROM users limit ?,?"
	// GetUserByUsernameSQL is SQL Clause to get user data by username
	GetUserByUsernameSQL = "SELECT id, username, first_name, last_name, email, date_joined, last_login, is_active, is_staff, is_superuser FROM users where username = ?"
)

// Repository define user repository methods interface
type Repository interface {
	// Get returns the user with the specified user ID.
	// Get(ctx context.Context, id string) (*model.User, error)
	List(ctx context.Context, id *ListUsersRequest) ([]model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

// repository persists user in database
type repository struct {
	db     *sqlx.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *sqlx.DB, log log.Logger) Repository {
	return repository{db: db, logger: log}
}

// List return slice of users base on request parameters
func (r repository) List(ctx context.Context, list *ListUsersRequest) ([]model.User, error) {
	users := make([]model.User, 0)
	offset := (list.Page - 1) * list.Limit
	err := r.db.Select(&users, ListUsersSQL, offset, list.Limit)
	if err != nil {
		r.logger.Errorf("Failed to select users %s", err)
		return nil, err
	}
	return users, nil
}

// GetByUsername user from username
func (r repository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, GetUserByUsernameSQL, username)
	if err != nil {
		r.logger.Errorf("Failed to select users %s", err)
		return nil, err
	}
	return &user, nil
}
