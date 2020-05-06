package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/qreasio/go-starter-kit/pkg/model"
)

// Service represents user application interface
type Service interface {
	ListUsers(ctx context.Context, listRequest *ListUsersRequest) ([]model.User, error)
}

type userService struct {
	repo      Repository
	validator *validator.Validate
}

// NewService return new userService instance
func NewService(repo Repository, validate *validator.Validate) Service {
	return userService{repo: repo, validator: validate}
}

// ListUsers implement service for list users
func (s userService) ListUsers(ctx context.Context, req *ListUsersRequest) ([]model.User, error) {
	if err := req.Validate(s.validator); err != nil {
		return []model.User{}, err
	}
	return s.repo.List(ctx, req)
}

// ListUsersRequest for representing request parameters on list users endpoint
type ListUsersRequest struct {
	model.Pagination
	Search string
}

// Validate validates the ListUsersRequest fields.
func (listRequest ListUsersRequest) Validate(validator *validator.Validate) error {
	return validator.Struct(listRequest)
}
