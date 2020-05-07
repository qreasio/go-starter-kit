package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository_go "github.com/qreasio/go-starter-kit/internal/user"
	model "github.com/qreasio/go-starter-kit/pkg/model"
)

// Repository is a mock of Repository interface
type Repository struct {
	ctrl     *gomock.Controller
	recorder *RepositoryMockRecorder
}

// RepositoryMockRecorder is the mock recorder for Repository
type RepositoryMockRecorder struct {
	mock *Repository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *Repository {
	mock := &Repository{ctrl: ctrl}
	mock.recorder = &RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Repository) EXPECT() *RepositoryMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *Repository) List(ctx context.Context, id *repository_go.ListUsersRequest) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, id)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *RepositoryMockRecorder) List(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Repository)(nil).List), ctx, id)
}
