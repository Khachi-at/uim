package uim

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/uim/wire/pkt"
)

type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock: mock}
	return mock
}

func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

func (m *MockDispatcher) Push(gateway string, channels []string, p *pkt.LogicPkt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", gateway, channels, p)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDispatcherMockRecorder) Push(gateway, channels, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockDispatcher)(nil).Push), gateway, channels, p)
}
