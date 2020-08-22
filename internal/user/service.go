package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/qreasio/go-starter-kit/pkg/model"
)

// Service represents user application interface
type Service interface {
	ListUsers(ctx context.Context, listRequest *ListUsersRequest) ([]model.User, error)
}

type userService struct {
	repo      Repository
	validator *validator.Validate
	logger    log.Logger
}

// NewService return new userService instance
func NewService(repo Repository, validate *validator.Validate, logger log.Logger) Service {
	return userService{repo: repo, validator: validate, logger: logger}
}

// ListUsers implement service for list users
func (s userService) ListUsers(ctx context.Context, req *ListUsersRequest) ([]model.User, error) {
	if err := req.Validate(s.validator); err != nil {
		s.logger.Errorf("Request validation failed on list users service : %s", err.Error())
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

// NewListUsersRequest construct new ListUsersRequest
func NewListUsersRequest() ListUsersRequest {
	return ListUsersRequest{
		Pagination: *model.NewPagination(),
		Search:     "",
	}
}
