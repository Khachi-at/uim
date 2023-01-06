package uim

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

func NewMockService(crtl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: crtl}
	mock.recorder = &MockServiceMockRecorder{mock: mock}
	return mock
}

// EXPECT return s an object tha allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetMeta mocks base method.
func (m *MockService) GetMeta() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeta")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetMeta indicates an expected call of GetMeta.
func (mr *MockServiceMockRecorder) GetMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeta", reflect.TypeOf((*MockService)(nil).GetMeta()))
}

// ServiceID mocks base method.
func (m *MockService) ServiceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceID")
	ret0, _ := ret[0].(string)
	return ret0
}
