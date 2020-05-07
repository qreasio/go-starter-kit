package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/qreasio/go-starter-kit/internal/user"
	model "github.com/qreasio/go-starter-kit/pkg/model"
)

// Service is a mock of Service interface
type Service struct {
	ctrl     *gomock.Controller
	recorder *ServiceMockRecorder
}

// ServiceMockRecorder is the mock recorder for Service
type ServiceMockRecorder struct {
	mock *Service
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *Service {
	mock := &Service{ctrl: ctrl}
	mock.recorder = &ServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Service) EXPECT() *ServiceMockRecorder {
	return m.recorder
}

// ListUsers mocks base method
func (m *Service) ListUsers(ctx context.Context, listRequest *user.ListUsersRequest) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, listRequest)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *ServiceMockRecorder) ListUsers(ctx, listRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*Service)(nil).ListUsers), ctx, listRequest)
}
